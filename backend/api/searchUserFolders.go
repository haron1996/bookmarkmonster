package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func SearchUserFolders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	folderName := chi.URLParam(r, "folderName")

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

	pool.Config().MinConns = 50

	pool.Config().MaxConns = 150

	pool.Config().MaxConnIdleTime = 5 * time.Second

	defer pool.Close()

	q := sqlc.New(pool)

	percent := "%"

	folderName = fmt.Sprintf("%s%s%s", percent, folderName, percent)

	params := sqlc.SearchUserFoldersParams{
		FolderName: folderName,
		UserID:     payload.UserID,
	}

	folders, err := q.SearchUserFolders(ctx, params)
	if err != nil {
		log.Printf("coult not get folders: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	utils.JsonResponse(w, folders)
}
