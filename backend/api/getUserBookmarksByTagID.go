package api

import (
	"errors"
	"log"
	"net/http"
	"sort"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func (h *BaseHandler) GetUserBookmarksByTagID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tagid := chi.URLParam(r, "tagid")

	// const pLoad mw.ContextKey = "payload"

	// payload := ctx.Value(pLoad).(*token.PayLoad)

	// userid := payload.UserID

	q := sqlc.New(h.db)

	bookmarkTags, err := q.GetUserBookmarksByTagID(ctx, tagid)
	if err != nil {
		var pgErr *pgconn.PgError

		switch {
		case errors.As(err, &pgErr):
			log.Printf("could not get bookmarks by tag id: %v", pgErr)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		default:
			log.Printf("could not get bookmarks by tag id: %v", err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}

	var bookmarks []sqlc.Bookmark

	for _, b := range bookmarkTags {
		bookmark, err := q.GetBookmarkByID(ctx, b.BookmarkID)
		if err != nil {
			var pgErr *pgconn.PgError

			switch {
			case errors.As(err, &pgErr):
				log.Printf("could not get bookmarks by tag id: %v", pgErr)
				utils.Response(w, "something went wrong", http.StatusInternalServerError)
				return
			case errors.Is(err, pgx.ErrNoRows):
				log.Println("bookmark by id not found")
				utils.Response(w, "bookmark by id not found", http.StatusNoContent)
				return
			default:
				log.Printf("could not get bookmarks by tag id: %v", err)
				utils.Response(w, "something went wrong", http.StatusInternalServerError)
				return
			}
		}

		bookmarks = append(bookmarks, bookmark)
	}

	sort.Slice(bookmarks, func(i, j int) bool {
		return bookmarks[i].Added.After(bookmarks[j].Added)
	})

	utils.JsonResponse(w, bookmarks)
}
