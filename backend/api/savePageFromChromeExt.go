package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type savePageFromChromeExtensionRequest struct {
	URL     string        `json:"url"`
	Title   string        `json:"title"`
	Tags    []sqlc.Tag    `json:"tags"`
	Folders []sqlc.Folder `json:"folders"`
}

func SavePageFromChromeExtension(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *savePageFromChromeExtensionRequest

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
	host, standardURL, err := getHostAndStandardURLFromURL(req.URL)
	if err != nil {
		log.Printf("could not get host and standard url from req url: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*token.PayLoad)

	log.Printf("payload: %v", payload)

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

	createBookmarkParams := sqlc.CreateBookmarkParams{}

	if len(req.Folders) >= 1 {
		for _, folder := range req.Folders {
			createBookmarkParams = sqlc.CreateBookmarkParams{
				ID:       utils.RandomString(),
				Title:    req.Title,
				Bookmark: standardURL,
				Host:     host,
				UserID:   payload.UserID,
				Notes:    pgtype.Text{},
				FolderID: pgtype.Text{String: folder.FolderID, Valid: true},
			}
		}
	} else {
		createBookmarkParams = sqlc.CreateBookmarkParams{
			ID:       utils.RandomString(),
			Title:    req.Title,
			Bookmark: standardURL,
			Host:     host,
			UserID:   payload.UserID,
			Notes:    pgtype.Text{},
			FolderID: pgtype.Text{},
		}
	}

	bookmark, err := q.CreateBookmark(ctx, createBookmarkParams)
	if err != nil {
		log.Printf("could not save page from chrome extension: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	if len(req.Tags) >= 1 {
		for _, tag := range req.Tags {
			tagBookmarkParams := sqlc.TagBookmarkParams{
				BookmarkID: bookmark.ID,
				TagID:      tag.ID,
			}

			_, err := q.TagBookmark(ctx, tagBookmarkParams)
			if err != nil {
				log.Printf("failed to tag bookmark saved from chrome extension: %v", err)
				utils.Response(w, "something went wrong", 500)
				return
			}
		}
	}

	utils.JsonResponse(w, bookmark)
}

func getHostAndStandardURLFromURL(URL string) (string, string, error) {
	var host string

	var normalizedURL string

	if strings.Contains(URL, "?") {
		u, err := url.ParseRequestURI(URL)
		if err != nil {
			return "", "", err
		}

		if u.Scheme == "https" {

			host = u.Host

			normalizedURL = fmt.Sprintf(`%v`, u)
		} else {
			return "", "", err
		}

	} else {
		parsedUrl, err := url.Parse(URL)
		if err != nil {
			return "", "", err
		}

		if parsedUrl.Scheme == "https" {
			host = parsedUrl.Host

			normalizedURL = URL
		} else {
			host = parsedUrl.String()

			normalizedURL = fmt.Sprintf(`https://%s`, URL)
		}

	}

	host, _ = strings.CutPrefix(host, "www.")

	return host, normalizedURL, nil
}
