package api

import (
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func GetUserRecentBookmarks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

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

	bookmarks, err := q.GetUserBookmarks(ctx, payload.UserID)
	if err != nil {
		log.Printf("could not get user bookmarks: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	var recentUserBookmarks []sqlc.Bookmark

	for _, b := range bookmarks {
		timeSinceAdded := time.Now().UTC().Sub(b.Added).Hours()

		week := 168 * time.Hour

		if timeSinceAdded < week.Hours() {
			recentUserBookmarks = append(recentUserBookmarks, b)
		}
	}

	utils.JsonResponse(w, recentUserBookmarks)
}
