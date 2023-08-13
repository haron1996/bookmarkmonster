package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/kwandapchumba/bookmarkmonster/api"
	connection "github.com/kwandapchumba/bookmarkmonster/db"
	"github.com/kwandapchumba/bookmarkmonster/mw"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.AllowContentEncoding("application/json", "application/x-www-form-urlencoded"))
	r.Use(middleware.CleanPath)
	r.Use(middleware.RedirectSlashes)

	h := api.NewBaseHandler(connection.ConnectDB())

	r.Route("/auth", func(r chi.Router) {
		r.Get("/get-google-login-url", h.GetGoogleAuthUrl)
		r.Post("/register-google-user", h.RegisterGoogleUser)
	})

	r.Route("/authenticated", func(r chi.Router) {
		r.Use(mw.AuthenticateRequest())
		r.Route("/tags", func(r chi.Router) {
			r.Post("/create-tag", h.CreateTag)
		})
	})

	return r
}
