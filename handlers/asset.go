package handlers

import (
	"AssetTrack/models"
	services "AssetTrack/service"
	"AssetTrack/utils"
	"net/http"
)

func CreateAsset(w http.ResponseWriter, r *http.Request) {
	var body models.AssetRequest

	if err := utils.ParseBody(r, &body); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "failed to parse request body")
		return
	}

	response := services.CreateAsset(body)

	if response.Err != nil {
		utils.RespondError(w, response.StatusCode, response.Err, response.Message)
		return
	}

	utils.RespondJSON(w, response.StatusCode, struct {
		Message string `json:"message"`
	}{
		Message: "asset created successfully",
	})
}

func GetAssets(w http.ResponseWriter, r *http.Request) {

	response := services.GetAssets()

	if response.Err != nil {
		utils.RespondError(
			w,
			response.StatusCode,
			response.Err,
			"failed to fetch assets",
		)
		return
	}

	utils.RespondJSON(w, response.StatusCode, struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Message: "assets fetched successfully",
		Data:    response.Data,
	})
}
