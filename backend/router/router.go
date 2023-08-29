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
		AllowedOrigins:   []string{"http://localhost:5173", "https://bookmarkmonster.xyz"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH", "HEAD"},
		AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
		Debug:            false,
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.AllowContentEncoding("application/json", "application/x-www-form-urlencoded"))
	r.Use(middleware.CleanPath)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.RealIP)

	h := api.NewBaseHandler(connection.ConnectDB())

	r.Route("/auth", func(r chi.Router) {
		r.Get("/get-google-login-url", h.GetGoogleAuthUrl)
		r.Post("/register-google-user", h.RegisterGoogleUser)
	})

	r.Route("/authenticated", func(r chi.Router) {
		r.Use(mw.AuthenticateRequest())
		r.Route("/tags", func(r chi.Router) {
			r.Get("/", h.GetAllUserTags)
			r.Get("/{tagName}", h.SearchUserTags)
			r.Post("/create-tag", h.CreateTag)
			r.Patch("/renameTag", h.RenameTag)
			r.Patch("/trashTag", h.TrashTag)
			r.Patch("/restoreTag", h.RestoreTag)
		})

		r.Route("/bookmarks", func(r chi.Router) {
			r.Get("/", h.GetUserBookmarks)
			r.Get("/{tagid}", h.GetUserBookmarksByTagID)
			r.Get("/search/{title}", h.SearchBookmarks)
			r.Get("/bookmarkTags/{bookmarkid}", h.GetBookmarkTags)
			r.Post("/add", h.AddBookmark)
			r.Patch("/tagBookmark", h.TagBookmark)
			r.Patch("/renameBookmark", h.RenameBookmark)
			r.Patch("/trashBookmarks", h.TrashBookmark)
			r.Delete("/deleteTagFromBookmark", h.DeleteTagFromBookmark)
		})
	})

	r.Route("/admin", func(r chi.Router) {
		r.Get("/users/{code}", h.GetAllUsers)
	})

	return r
}
