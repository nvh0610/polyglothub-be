package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	customStatus "learn/internal/common/error"
	"learn/internal/controller"
	"learn/internal/repository"
	"learn/job/schedule"
	"learn/pkg/logger"
	mdw "learn/pkg/middleware"
	"learn/pkg/resp"
	"learn/platform/mysqldb"
	"learn/platform/redisdb"
	"net/http"
)

func InitRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/ping", func(writer http.ResponseWriter, request *http.Request) {
		resp.Return(writer, http.StatusOK, customStatus.SUCCESS, "success")
	})

	schedule.Start()
	mysqlConn, err := mysqldb.NewMysqlConnection()
	if err != nil {
		logger.Error(err.Error())
	}

	redisConn, err := redisdb.NewRedisConnection()
	if err != nil {
		logger.Error(err.Error())
	}

	baseRepo := repository.NewRegistryRepo(mysqlConn)
	baseController := controller.NewRegistryController(baseRepo, redisConn)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", baseController.AuthCtrl.Login)
		r.Post("/forget-password", baseController.AuthCtrl.ForgetPassword)
		r.Post("/verify-otp", baseController.AuthCtrl.VerifyOtp)
		r.Post("/reset-password", baseController.AuthCtrl.ResetPassword)
		r.With(mdw.JwtMiddleware).Post("/change-password", baseController.AuthCtrl.ChangePassword)
	})

	r.Route("/vocabulary", func(r chi.Router) {
		r.Use(mdw.JwtMiddleware)
		r.Get("/", baseController.VocabularyCtrl.ListVocabulary)
		r.Post("/", baseController.VocabularyCtrl.CreateVocabulary)
		r.Put("/{id}", baseController.VocabularyCtrl.UpdateVocabulary)
		r.Delete("/{id}", baseController.VocabularyCtrl.DeleteVocabulary)
	})

	r.Route("/category", func(r chi.Router) {
		r.Use(mdw.JwtMiddleware)
		r.Get("/", baseController.CategoryCtrl.ListCategory)
		r.Post("/", baseController.CategoryCtrl.CreateCategory)
		r.Put("/{id}", baseController.CategoryCtrl.UpdateCategory)
		r.Delete("/{id}", baseController.CategoryCtrl.DeleteCategory)
	})

	r.Route("/user", func(r chi.Router) {
		r.Use(mdw.JwtMiddleware)
		r.Get("/{id}", baseController.UserCtrl.GetUserById)
		r.Post("/", baseController.UserCtrl.CreateUser)
		r.Put("/{id}", baseController.UserCtrl.UpdateUser)
		r.Delete("/{id}", baseController.UserCtrl.DeleteUser)
		r.Get("/", baseController.UserCtrl.ListUser)
		r.Post("/update-role", baseController.UserCtrl.UpdateRole)
	})

	r.Route("/flashcard-daily", func(r chi.Router) {
		r.Use(mdw.JwtMiddleware)
		r.Get("/", baseController.FlashCardDailyCtrl.GetFlashCardDaily)
		r.Post("/confirm", baseController.FlashCardDailyCtrl.ConfirmFlashCardDaily)
	})

	return r
}
