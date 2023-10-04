package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type syncChromeBookmarksRequest struct {
	Bookmarks []sqlc.Bookmark `json:"bookmarks"`
}

func SyncChromeBookmarks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *syncChromeBookmarksRequest

	if err := jsonDecoder.Decode(&req); err != nil {
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		} else {
			log.Printf("error decoding request body to struct: %v", err)
			utils.Response(w, "something went wrong", http.StatusBadRequest)
			return
		}
	}

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*token.PayLoad)

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

	for _, bookmark := range req.Bookmarks {
		_, standardURL, err := GetHostAndStandardURLFromURL(bookmark.Bookmark)
		if err != nil {
			log.Printf("could not get host and standard url from req url: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}

		_, err = q.GetChromeBookmark(ctx, sqlc.GetChromeBookmarkParams{
			Bookmark: standardURL,
			Title:    bookmark.Title,
			UserID:   payload.UserID,
		})
		if err != nil {
			switch {
			case errors.Is(err, pgx.ErrNoRows):
				_, err := createBookmark(q, bookmark, ctx, payload)
				if err != nil {
					log.Printf("could not create bookmark: %v", err)
					utils.Response(w, "something went wrong", 500)
					return
				}

				continue
			default:
				log.Printf("could not get bookmark: %v", err)
				utils.Response(w, "something went wrong", 500)
				return
			}
		}

	}

	utils.JsonResponse(w, "sync succeeded")
}

func createBookmark(q *sqlc.Queries, bookmark sqlc.Bookmark, ctx context.Context, payload *token.PayLoad) (sqlc.Bookmark, error) {
	host, standardURL, err := GetHostAndStandardURLFromURL(bookmark.Bookmark)
	if err != nil {
		return sqlc.Bookmark{}, err
	}

	b, err := q.CreateBookmark(ctx, sqlc.CreateBookmarkParams{
		ID:         utils.RandomString(),
		Title:      bookmark.Title,
		Bookmark:   standardURL,
		Host:       host,
		UserID:     payload.UserID,
		Notes:      pgtype.Text{},
		FolderID:   pgtype.Text{},
		Fromchrome: true,
	})

	if err != nil {
		return sqlc.Bookmark{}, err
	}

	return b, nil
}
