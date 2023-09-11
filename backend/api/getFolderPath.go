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

func GetFolderPath(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

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

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*token.PayLoad)

	q := sqlc.New(pool)

	folderID := chi.URLParam(r, "folderID")

	f, err := q.GetFolder(ctx, sqlc.GetFolderParams{FolderID: folderID, UserID: payload.UserID})
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			log.Printf("folder not found: %v", err)
			utils.Response(w, "folder not found", http.StatusNotFound)
			return
		default:
			log.Printf("could not get folder: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}
	}

	fs, err := q.GetFolderPath(ctx, f.Label)
	if err != nil {
		log.Printf("could not get folder path: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	utils.JsonResponse(w, fs)
}
