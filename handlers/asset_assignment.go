package handlers

import (
	"AssetTrack/models"
	services "AssetTrack/service"
	"AssetTrack/utils"
	"net/http"
)

func AssignAsset(w http.ResponseWriter, r *http.Request) {

	var body models.AssignAssetRequest

	if parseErr := utils.ParseBody(r, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	response := services.AssignAsset(body)

	utils.RespondJSON(w, response.StatusCode, struct {
		Message string `json:"message"`
	}{
		Message: "asset assigned successfully",
	})
}

func ReturnAsset(w http.ResponseWriter, r *http.Request) {

	var body models.AssignAssetRequest

	if parseErr := utils.ParseBody(r, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	response := services.ReturnAsset(body)

	utils.RespondJSON(w, response.StatusCode, response)
}
