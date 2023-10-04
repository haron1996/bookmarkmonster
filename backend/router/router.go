package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/kwandapchumba/bookmarkmonster/api"
	cookieapis "github.com/kwandapchumba/bookmarkmonster/api/cookieApis"
	"github.com/kwandapchumba/bookmarkmonster/mw"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "https://bookmarkmonster.xyz", "https://*.bookmarkmonster.xyz"},
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

	// h := api.NewBaseHandler(connection.ConnectDB())

	r.Route("/auth", func(r chi.Router) {
		r.Get("/get-google-login-url", api.GetGoogleAuthUrl)
		r.Post("/register-google-user", api.RegisterGoogleUser)
		r.Post("/requestEmailVerificationCode", api.RequestEmailVerification)
		r.Post("/resendEmailConfirmationLink", api.ResendConfirmationLink)
		r.Post("/createUser", api.CreateUser)
		r.Post("/logUserIn", api.LogUserIn)
	})

	r.Route("/authenticated", func(r chi.Router) {
		r.Use(mw.AuthenticateRequest())

		r.Post("/requestPocketAuthorizationCode", api.RequestPocketAuthorizationCode)
		r.Post("/convertPocketRequestTokenToAccessToken", api.ConverPocketRequestTokenToAccessToken)
		r.Post("/importFromPocket", api.ImportFromPocket)

		r.Route("/tags", func(r chi.Router) {
			r.Get("/", api.GetAllUserTags)
			r.Get("/{tagName}", api.SearchUserTags)
			r.Post("/create-tag", api.CreateTag)
			r.Patch("/renameTag", api.RenameTag)
			r.Patch("/trashTag", api.TrashTag)
			r.Patch("/restoreTag", api.RestoreTag)
			r.Get("/searchTagBookmarks/{tag_id}/{query}", api.SearchTagBookmarks)
		})

		r.Route("/bookmarks", func(r chi.Router) {
			r.Get("/", api.GetUserBookmarks)
			r.Get("/{tagid}", api.GetUserBookmarksByTagID)
			r.Get("/search/{title}", api.SearchBookmarks)
			r.Get("/bookmarkTags/{bookmarkid}", api.GetBookmarkTags)
			r.Post("/add", api.AddBookmark)
			r.Patch("/tagBookmark", api.TagBookmark)
			r.Patch("/renameBookmark", api.RenameBookmark)
			r.Patch("/trashBookmarks", api.TrashBookmark)
			r.Delete("/deleteTagFromBookmark", api.DeleteTagFromBookmark)
			r.Get("/getUserRecentBookmarks", api.GetUserRecentBookmarks)

			// lates
			r.Post("/createBookmark", api.CreateRootBookmark)
			r.Get("/getRootBookmarks", api.GetUserRootBookmarks)
			r.Get("/getAllUserBookmarks", api.GetAllBookmarks)
			r.Get("/getBookmarksInTrash", api.GetBookmarksInTrash)
			r.Patch("/updateBookmark", api.UpdateBookmark)
			r.Patch("/moveBookmarks", api.MoveBookmarks)
			r.Patch("/updateBookmarkHtml", api.UpdateBookmarkHtml)
			r.Post("/captureFullpageScreeshot", api.CaptureFullpageScreeshot)
			r.Post("/captureAboveTheFoldScreenshot", api.GetAboveTheFoldScreenshot)
			r.Get("/getFullpageScreenshots", api.GetFullpageScreenshots)
			r.Get("/getAboveFoldScreenshots", api.GetAboveFoldScreenshots)
			r.Post("/createAndTagBookmark", api.CreateAndTagBookmark)
			r.Post("/savePageFromChromeExtension", api.SavePageFromChromeExtension)
			r.Post("/syncChromeBookmarks", api.SyncChromeBookmarks)
		})

		r.Route("/collections", func(r chi.Router) {
			r.Post("/createCollection", api.CreateFolder)
			r.Get("/getUserRootCollections", api.GetUserRootFolders)
			r.Get("/getUserFoldersInRecycleBin", api.GetFoldersInTrash)
			r.Get("/getCollectionPath/{folderID}", api.GetFolderPath)
			r.Get("/getFolderBookmarks/{folderID}", api.GetFolderBookmarks)
			r.Get("/getFolderSubfolders/{folderID}", api.GetFolderSubfolders)
			r.Get("/getFolderPath/{folderID}", api.GetFolderPath)
			r.Patch("/updateCollection", api.UpdateFolder)
			r.Patch("/moveCollectionsToAnother", api.MoveCollectionsToAnother)
			r.Patch("/moveCollectionsToRoot", api.MoveCollectionsToRoot)
			r.Get("/getFolder/{folderID}", api.GetFolder)
			r.Get("/searchUserFolders/{folderName}", api.SearchUserFolders)
		})
	})

	r.Route("/admin", func(r chi.Router) {
		r.Get("/users/{code}", api.GetAllUsers)
		r.Get("/getWaitList/{code}", api.GetWaitList)
		r.Get("/readSupportMessages/{code}", api.ReadSupportMessages)
	})

	r.Route("/care", func(r chi.Router) {
		r.Post("/joinWaitlist", api.JoinWaitList)
		r.Post("/contactSupport", api.ContactSupport)
	})

	r.Route("/cookie", func(r chi.Router) {
		r.Get("/searchTags/{tagName}", cookieapis.SearchTags)
	})

	return r
}
