package handler

import (
	"RedMist/internal/app/handler/custom"
	"RedMist/internal/app/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	handler.UserHandler
	DB *gorm.DB
}

func (h *Handler) InitRoutes() *chi.Mux {
	userService := services.NewUserService(h.DB)
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	router.Use(middleware.Logger)
	router.Use(middleware.Heartbeat("/ping"))
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {})
	router.Route("/auth", func(auth chi.Router) {
		auth.Post("/register", h.Register.Register)
		auth.Post("/login", h.Login.Login)
	})
	router.Route("/", func(r chi.Router) {

	})

	return router
}
