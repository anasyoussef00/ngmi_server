package health_check

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterHandlers(r *chi.Mux, version string) {
	r.Get("/health-check", healthCheck(version))
	r.Head("/health-check", healthCheck(version))
}

func healthCheck(version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(fmt.Sprintf("OK: %s", version))); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}
