package auth

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	customStatus "learn/internal/common/error"
	"learn/internal/repository"
	"learn/internal/router/payload/request"
	"learn/internal/router/payload/response"
	"learn/pkg/jwt"
	"learn/pkg/password"
	"learn/pkg/resp"
	"learn/pkg/utils"
	"learn/platform/send_otp"
	"net/http"
	"time"
)

const (
	OTP_KEY          = "learn-language-%d"
	VALIDATE_OTP_KEY = "learn-language-validate-otp-%d"
)

type AuthController struct {
	repo      repository.Registry
	redis     *redis.Client
	sendEmail *send_otp.SendOtpEmail
}

func NewAuthController(userRepo repository.Registry, redis *redis.Client) Controller {
	return &AuthController{
		repo:      userRepo,
		redis:     redis,
		sendEmail: send_otp.NewSendOtpEmail(),
	}
}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	req := &request.LoginUserRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	user, err := a.repo.User().GetByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.USER_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if password.CheckPassword(req.Password, user.Password) != nil {
		resp.Return(w, http.StatusUnauthorized, customStatus.WRONG_PASSWORD, nil)
		return
	}

	token, err := jwt.GenerateToken(user.Id, user.Role)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, response.LoginResponse{
		AccessToken: token,
	})
}

func (a *AuthController) ChangePassword(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user_id").(int)

	req := &request.ChangePasswordRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	user, err := a.repo.User().GetById(userId)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if password.CheckPassword(req.OldPassword, user.Password) != nil {
		resp.Return(w, http.StatusUnauthorized, customStatus.WRONG_PASSWORD, nil)
		return
	}

	hashedPassword, err := password.HashPassword(req.NewPassword)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	user.Password = hashedPassword
	err = a.repo.User().Update(user)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (a *AuthController) ForgetPassword(w http.ResponseWriter, r *http.Request) {
	req := &request.ForgetPasswordRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	user, err := a.repo.User().GetByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.USER_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	randomOtp := utils.EncodeToString()
	err = a.redis.Set(r.Context(), fmt.Sprintf(OTP_KEY, user.Id), randomOtp, 3*time.Minute).Err()
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	go a.sendEmail.SendOtp(user.Username, randomOtp)
	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (a *AuthController) VerifyOtp(w http.ResponseWriter, r *http.Request) {
	req := &request.VerifyOtpRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	user, err := a.repo.User().GetByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.USER_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	exist, err := a.redis.Exists(r.Context(), fmt.Sprintf(OTP_KEY, user.Id)).Result()
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if exist != 1 {
		resp.Return(w, http.StatusForbidden, customStatus.FORBIDDEN, nil)
		return
	}

	otp, err := a.redis.Get(r.Context(), fmt.Sprintf(OTP_KEY, user.Id)).Result()
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if req.Otp != otp {
		resp.Return(w, http.StatusBadRequest, customStatus.WRONG_OTP, nil)
		return
	}

	if err = a.redis.Del(r.Context(), fmt.Sprintf(OTP_KEY, user.Id)).Err(); err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if err = a.redis.Set(r.Context(), fmt.Sprintf(VALIDATE_OTP_KEY, user.Id), 1, 3*time.Minute).Err(); err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}

func (a *AuthController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	req := &request.ResetPasswordRequest{}
	if err := utils.BindAndValidate(r, req); err != nil {
		resp.Return(w, http.StatusBadRequest, customStatus.INVALID_PARAMS, err.Error())
		return
	}

	user, err := a.repo.User().GetByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp.Return(w, http.StatusNotFound, customStatus.USER_NOT_FOUND, nil)
			return
		}

		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	exist, err := a.redis.Exists(r.Context(), fmt.Sprintf(VALIDATE_OTP_KEY, user.Id)).Result()
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if exist != 1 {
		resp.Return(w, http.StatusForbidden, customStatus.FORBIDDEN, nil)
		return
	}

	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	user.Password = hashedPassword
	err = a.repo.User().Update(user)
	if err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	if err = a.redis.Del(r.Context(), fmt.Sprintf(VALIDATE_OTP_KEY, user.Id)).Err(); err != nil {
		resp.Return(w, http.StatusInternalServerError, customStatus.INTERNAL_SERVER, err.Error())
		return
	}

	resp.Return(w, http.StatusOK, customStatus.SUCCESS, nil)
}
