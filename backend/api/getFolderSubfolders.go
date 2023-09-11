package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func GetFolderSubfolders(w http.ResponseWriter, r *http.Request) {
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

	folderID := chi.URLParam(r, "folderID")

	f, err := q.GetFolder(ctx, sqlc.GetFolderParams{FolderID: folderID, UserID: payload.UserID})
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			log.Printf("parent folder with id: %s not found", folderID)
			utils.Response(w, "parent folder not found", http.StatusNotFound)
			return
		default:
			log.Printf("could not get parent folder: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}
	}

	if f.UserID != payload.UserID {
		log.Printf("folder user id: %s and user id: %s dont match", f.UserID, payload.UserID)
		utils.Response(w, "folder userID and payload userID don't match", http.StatusForbidden)
		return
	}

	fs, err := q.GetFolderSubFolders(ctx, sqlc.GetFolderSubFoldersParams{UserID: f.UserID, SubfolderOf: pgtype.Text{String: f.FolderID, Valid: true}})
	if err != nil {
		log.Printf("could not get folder subfolders: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
		// switch {
		// case errors.Is(err, pgx.ErrNoRows):
		// 	log.Printf("subfolders of: %s not found", f.FolderName)
		// 	var subFolders []*sqlc.Folder
		// 	utils.JsonResponse(w, subFolders)
		// 	return
		// default:
		// 	log.Printf("could not get subfoders of: %s with: %v", f.FolderName, err)
		// 	utils.Response(w, "something went wrong", 500)
		// 	return
		// }
	}

	utils.JsonResponse(w, fs)
}
