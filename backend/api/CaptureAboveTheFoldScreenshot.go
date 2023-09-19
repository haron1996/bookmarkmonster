package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/mw"
	token "github.com/kwandapchumba/bookmarkmonster/token"
	"github.com/kwandapchumba/bookmarkmonster/utils"
	"github.com/kwandapchumba/bookmarkmonster/vultr"
)

type captureAboveTheFoldScreenshotReq struct {
	BookmarkID string `json:"bookmark_id"`
}

func GetAboveTheFoldScreenshot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	jsonDecoder := json.NewDecoder(r.Body)

	jsonDecoder.DisallowUnknownFields()

	var req *captureAboveTheFoldScreenshotReq

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

	page := browser.MustPage(b.Bookmark).MustWaitLoad().MustWaitDOMStable().MustWindowMaximize().MustWindowFullscreen()

	defer browser.MustClose()

	body := page.MustElement("body")

	body.MustEval(`() => this.style.overflowY = 'hidden'`)

	if err := page.Emulate(devices.Clear); err != nil {
		log.Printf("could not emulate provided device: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	screenShotByte, err := page.Screenshot(false, &proto.PageCaptureScreenshot{
		Format: proto.PageCaptureScreenshotFormatPng,
	})
	if err != nil {
		log.Printf("could not get screenshot byte: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	f, err := os.Create("screenshot.png")
	if err != nil {
		log.Printf("could not create screenshot file: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	defer f.Close()

	_, err = f.Write(screenShotByte)
	if err != nil {
		log.Printf("could not write screenshot byte: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	url, err := vultr.UploadImage(f, "above-fold-screenshots")
	if err != nil {
		log.Printf("could not upload screenshot: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	const pLoad mw.ContextKey = "payload"

	payload := ctx.Value(pLoad).(*token.PayLoad)

	args := sqlc.CreateBookmarkScreenshotParams{
		ID:                 utils.RandomString(),
		ScreenshotLocation: url,
		BookmarkID:         req.BookmarkID,
		UserID:             payload.UserID,
		Fullpage:           false,
	}

	screenshot, err := q.CreateBookmarkScreenshot(ctx, args)
	if err != nil {
		log.Printf("could not create bookmark screenshot: %v", err)
		utils.Response(w, "something went wrong", 500)
		return
	}

	utils.JsonResponse(w, screenshot)
}
