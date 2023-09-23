package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

type createAndTagBookmarkRequest struct {
	Bookmark string   `json:"bookmark"`
	Tags     []string `json:"tags"`
}

func CreateAndTagBookmark(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *createAndTagBookmarkRequest

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

	var host string

	var urlToOpen string

	if strings.Contains(req.Bookmark, "?") {
		u, err := url.ParseRequestURI(req.Bookmark)
		if err != nil {
			log.Println(err)
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		if u.Scheme == "https" {

			host = u.Host

			urlToOpen = fmt.Sprintf(`%v`, u)
		} else {
			utils.Response(w, "invalid url", http.StatusBadRequest)
			return
		}

	} else {
		parsedUrl, err := url.Parse(req.Bookmark)
		if err != nil {
			utils.Response(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		if parsedUrl.Scheme == "https" {
			host = parsedUrl.Host

			urlToOpen = req.Bookmark
		} else {
			host = parsedUrl.String()

			urlToOpen = fmt.Sprintf(`https://%s`, req.Bookmark)
		}

	}

	host, _ = strings.CutPrefix(host, "www.")

	u := launcher.New().UserDataDir("~/.config/google-chrome").Leakless(true).NoSandbox(true).Headless(true).MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()

	page := browser.MustPage(urlToOpen).MustWaitLoad()

	defer browser.MustClose()

	pageHasTitle, title, err := page.Has("title")
	if err != nil {
		log.Printf("could not check if page has title: %v", err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	var pageTitle string

	if pageHasTitle {
		pageTitle = strings.TrimSpace(title.MustText())
	} else {
		pageTitle = req.Bookmark
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

	b, err := q.CreateBookmark(ctx, sqlc.CreateBookmarkParams{
		ID:       utils.RandomString(),
		Title:    pageTitle,
		Bookmark: urlToOpen,
		Host:     host,
		UserID:   payload.UserID,
		Notes:    pgtype.Text{},
		FolderID: pgtype.Text{},
	})
	if err != nil {
		log.Printf("could notcreate bookmark: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	for _, tag := range req.Tags {
		_, err := q.TagBookmark(ctx, sqlc.TagBookmarkParams{
			BookmarkID: b.ID,
			TagID:      tag,
		})
		if err != nil {
			log.Printf("could not tag bookmark: %v", err)
			utils.Response(w, "something went wrong", 500)
			return
		}
	}

	utils.JsonResponse(w, b)
}
