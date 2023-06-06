package authentication_routes

import (
	"github.com/go-chi/chi/v5"
	authenticationResource "ngmi_server/internal/resources/authentication_resource"
	authenticationService "ngmi_server/internal/services/authentication_service"
	"ngmi_server/pkg/log"
)

func RegisterHandlers(r chi.Router, service authenticationService.Service, logger log.Logger) {
	r.Route("/auth", func(r chi.Router) {
		res := authenticationResource.Resource{Service: service, Logger: logger}

		r.Post("/register", res.Register)
		r.Post("/login", res.Login)
	})
}
