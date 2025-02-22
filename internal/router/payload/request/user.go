package request

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	FullName string `json:"fullname" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	FullName string `json:"fullname" validate:"required"`
}

type UpdateRoleRequest struct {
	Role string `json:"role" validate:"required"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}

type ForgetPasswordRequest struct {
	Username string `json:"username" validate:"required"`
}

type VerifyOtpRequest struct {
	Username string `json:"username" validate:"required"`
	Otp      string `json:"otp" validate:"required"`
}

type ResetPasswordRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
