package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	customStatus "learn/internal/common/error"
	"learn/internal/controller"
	"learn/internal/repository"
	"learn/pkg/logger"
	"learn/pkg/resp"
	"learn/platform/mysqldb"
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

	mysqlConn, err := mysqldb.NewMysqlConnection()
	if err != nil {
		logger.Error(err.Error())
	}

	baseRepo := repository.NewRegistryRepo(mysqlConn)
	baseController := controller.NewRegistryController(baseRepo)
	r.Route("/user", func(r chi.Router) {
		r.Get("/{id}", baseController.UserCtrl.GetUserById)
		r.Post("/", baseController.UserCtrl.CreateUser)
	})

	return r
}
