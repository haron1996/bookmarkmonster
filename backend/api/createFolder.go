package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type createFolderRequest struct {
	Name           string `json:"name"`
	ParentFolderID string `json:"parent_folder_id"`
	Description    string `json:"description"`
}

func CreateFolder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *createFolderRequest

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

	var createFolderParams sqlc.CreateFolderParams

	if req.ParentFolderID == "" {
		exists, err := q.GetUserRootFolderByName(ctx, sqlc.GetUserRootFolderByNameParams{UserID: payload.UserID, FolderName: strings.ToLower(req.Name)})
		if err != nil {
			log.Printf("could not check if user root folder by name exists: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}

		if exists {
			log.Println("folder with name already exists in root")
			utils.Response(w, fmt.Sprintf("folder %s already exists here", req.Name), http.StatusConflict)
			return
		}

		label := utils.FolderLabel()

		createFolderParams = sqlc.CreateFolderParams{
			FolderID:          utils.RandomString(),
			FolderName:        req.Name,
			SubfolderOf:       pgtype.Text{String: "", Valid: false},
			UserID:            payload.UserID,
			Path:              label,
			Label:             label,
			FolderDescription: req.Description,
		}
	} else {
		parentFolder, err := q.GetFolder(ctx, sqlc.GetFolderParams{FolderID: req.ParentFolderID, UserID: payload.UserID})
		if err != nil {
			switch {
			case errors.Is(err, pgx.ErrNoRows):
				log.Printf("folder with id: %s not found", req.ParentFolderID)
				utils.Response(w, fmt.Sprintf("folder with id: %s not found", req.ParentFolderID), http.StatusNotFound)
				return
			default:
				log.Printf("parent folder with id: %s not found: %v", req.ParentFolderID, err)
				utils.Response(w, "something went wrong", 500)
				return
			}
		}

		label := utils.FolderLabel()

		path := strings.Join([]string{parentFolder.Path, label}, ".")

		createFolderParams = sqlc.CreateFolderParams{
			FolderID:          utils.RandomString(),
			FolderName:        req.Name,
			SubfolderOf:       pgtype.Text{String: req.ParentFolderID, Valid: true},
			UserID:            payload.UserID,
			Path:              path,
			Label:             label,
			FolderDescription: req.Description,
		}
	}

	folder, err := q.CreateFolder(ctx, createFolderParams)
	if err != nil {
		var pgErr *pgconn.PgError

		switch {
		case errors.As(err, &pgErr):
			switch {
			case pgErr.Error() == `ERROR: duplicate key value violates unique constraint "folder_subfolder_of_folder_name_key" (SQLSTATE 23505)`:
				log.Printf("could not create sub folder: %v", pgErr.Error())
				utils.Response(w, fmt.Sprintf("folder %s already exists here", req.Name), http.StatusConflict)
				return
			default:
				log.Printf("could not create folder with pgErr: %v", pgErr)
				utils.Response(w, "something went wrong", http.StatusInternalServerError)
				return
			}
		default:
			log.Printf("could not create folder: %v", err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	}

	utils.JsonResponse(w, folder)

}
