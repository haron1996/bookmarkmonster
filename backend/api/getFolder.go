package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func GetFolder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	folderID := chi.URLParam(r, "folderID")

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

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*token.PayLoad)

	args := sqlc.GetFolderParams{
		FolderID: folderID,
		UserID:   payload.UserID,
	}

	folder, err := q.GetFolder(ctx, args)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			log.Println("folder not found")
			utils.Response(w, "folder not found", http.StatusNotFound)
			return
		default:
			log.Printf("could not get folder: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}
	}

	utils.JsonResponse(w, folder)
}
