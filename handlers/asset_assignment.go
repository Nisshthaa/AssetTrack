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

	err, statusCode, message := services.AssignAsset(body)
	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to assign assets")
		return
	}

	utils.RespondJSON(w, statusCode, struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
}

func ReturnAsset(w http.ResponseWriter, r *http.Request) {

	var body models.AssignAssetRequest

	if parseErr := utils.ParseBody(r, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	err, statusCode, message := services.ReturnAsset(body)
	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to return assets")
		return
	}

	utils.RespondJSON(w, statusCode, struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
}

func GetAssetHistory(w http.ResponseWriter, r *http.Request) {

	assetID := r.PathValue("assetID")

	history, statusCode, message, err := services.GetAssetHistory(assetID)
	if err != nil {
		utils.RespondError(w, statusCode, err, message)
		return
	}

	utils.RespondJSON(w, statusCode, struct {
		Message string      `json:"message"`
		History interface{} `json:"history"`
	}{
		Message: message,
		History: history,
	})
}
