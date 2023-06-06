package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	authentication_repository "ngmi_server/internal/repositories/authentication_repository"
	"ngmi_server/internal/routes/authentication_routes"
	authentication_service "ngmi_server/internal/services/authentication_service"
	"ngmi_server/pkg/db"
	"ngmi_server/pkg/log"
	"time"
)

func BuildHandler(db *db.DB, logger log.Logger) *chi.Mux {
	r := chi.NewRouter()
	registerMiddlewares(r)
	registerHandlers(r, db, logger)
	return r
}

func registerMiddlewares(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
}

func registerHandlers(r *chi.Mux, db *db.DB, logger log.Logger) {
	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			authentication_routes.RegisterHandlers(r, authentication_service.NewService(authentication_repository.NewRepository(db, logger), logger), logger)
		})
	})
}
