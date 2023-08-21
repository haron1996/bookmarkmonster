package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	sqlc "github.com/kwandapchumba/bookmarkmonster/db/sqlc"
	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func (h *BaseHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	if code != config.AdminCode {
		log.Println("wrong code")
		utils.Response(w, "unauthorized", http.StatusForbidden)
		return
	}

	users, err := sqlc.New(h.pool).GetUsers(r.Context())
	if err != nil {
		log.Println(err)
		utils.Response(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, users)
}
