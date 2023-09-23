package api

import (
	"errors"
	"log"
	"net/http"
	"sort"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func GetUserBookmarksByTagID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tagid := chi.URLParam(r, "tagid")

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Println(err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	pool, err := pgxpool.New(ctx, config.DBString)
	if err != nil {
		log.Printf("could not create new pool: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	defer pool.Close()

	q := sqlc.New(pool)

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
				continue
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
