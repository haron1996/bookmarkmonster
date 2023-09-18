package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type getPageHtmlRequest struct {
	BookmarkID string `json:"bookmark_id"`
}

func UpdateBookmarkHtml(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *getPageHtmlRequest

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

	b, err := q.GetBookmarkByID(ctx, req.BookmarkID)
	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			log.Printf("bookmark not found: %v", err)
			utils.Response(w, "bookmark not found", http.StatusNotFound)
			return
		default:
			log.Printf("could not get bookmark: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}
	}

	u := launcher.New().UserDataDir("~/.config/google-chrome").Leakless(true).NoSandbox(true).Headless(true).MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	page := browser.MustPage(b.Bookmark).MustWaitLoad().MustWaitDOMStable()

	defer browser.MustClose()

	log.Println(page.MustHTML())
}
