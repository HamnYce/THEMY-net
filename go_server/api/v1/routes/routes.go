package apiv1routes

import (
	handlers "themynet/api/v1/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func SetupRoutes() (r chi.Router) {
	r = chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.AllowContentType("application/json"),
		middleware.CleanPath,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		cors.Handler(
			cors.Options{
				// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
				// AllowedOrigins: []string{"https://*", "http://*", "localhost:3000"},
				AllowedOrigins: []string{"https://localhost:3000", "http://localhost:3000"},
				// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: false,
				MaxAge:           300, // Maximum value not ignored by any of major browsers
			},
		),
	)

	r.Post("/CreateHosts", handlers.CreateHostsHandler)
	r.Post("/RetrieveHosts", handlers.RetrieveHostsHandler)
	r.Post("/UpdateHosts", handlers.UpdateHostsHandler)
	r.Post("/DeleteHosts", handlers.DeleteHostsHandler)

	return r
}
