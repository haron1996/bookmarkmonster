package api

import (
	"net/http"

	"github.com/kwandapchumba/bookmarkmonster/utils"
)

func (h *BaseHandler) GetGoogleAuthUrl(w http.ResponseWriter, r *http.Request) {

	authUrl, err := utils.GenerateGoogleAuthUrl()
	if err != nil {
		utils.Response(w, "could not generate auth url", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, authUrl)
}
