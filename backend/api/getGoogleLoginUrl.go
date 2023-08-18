package api

import (
	"log"
	"net/http"

	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func (h *BaseHandler) GetGoogleAuthUrl(w http.ResponseWriter, r *http.Request) {

	authUrl, err := utils.GenerateGoogleAuthUrl()
	if err != nil {
		log.Printf("failder to generate google auth url: %v", err)
		utils.Response(w, "could not generate auth url", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, authUrl)
}
