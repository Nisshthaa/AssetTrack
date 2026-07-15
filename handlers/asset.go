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

	statusCode, message, err := services.CreateAsset(body)

	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to create asset")
		return
	}

	utils.RespondJSON(w, statusCode, struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
}

func GetAssets(w http.ResponseWriter, r *http.Request) {

	assets, statusCode, message, err := services.GetAssets()

	if err != nil {
		utils.RespondError(w, statusCode, err, "failed to fetch assets")
		return
	}

	utils.RespondJSON(w, statusCode, struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Message: message,
		Data:    assets,
	})
}

func GetAssetByID(w http.ResponseWriter, r *http.Request) {

	assetID := r.PathValue("assetID")

	asset, statusCode, message, err := services.GetAssetByID(assetID)

	if err != nil {
		utils.RespondError(w, statusCode, err, message)
		return
	}

	utils.RespondJSON(w, statusCode, struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Message: "asset fetched successfully",
		Data:    asset,
	})
}

func UpdateAsset(w http.ResponseWriter, r *http.Request) {

	var body *models.UpdateAssetRequest

	if parseErr := utils.ParseBody(r, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	assetID := r.PathValue("assetID")

	statusCode, message, err := services.UpdateAsset(assetID, body)

	if err != nil {
		utils.RespondError(w, statusCode, err, message)
		return
	}

	utils.RespondJSON(w, statusCode, struct {
		Message string `json:"message"`
	}{
		Message: "asset updated successfully",
	})
}

func AssetSentToRepair(w http.ResponseWriter, r *http.Request) {

	assetID := r.PathValue("assetID")

	response := services.AssetSentToRepair(assetID)

	utils.RespondJSON(w, response.StatusCode, struct {
		Message string `json:"message"`
	}{
		Message: "asset repair request sent successfully",
	})
}

func AssetRepairCompleted(w http.ResponseWriter, r *http.Request) {

	assetID := r.PathValue("assetID")

	response := services.AssetRepairCompleted(assetID)

	utils.RespondJSON(w, response.StatusCode, struct {
		Message string `json:"message"`
	}{
		Message: "asset repair status updated successfully",
	})
}

func DeleteAsset(w http.ResponseWriter, r *http.Request) {

	assetID := r.PathValue("assetID")

	err, statusCode, message := services.DeleteAsset(assetID)

	if err != nil {
		utils.RespondError(w, statusCode, err, message)
		return
	}

	utils.RespondJSON(w, statusCode, struct {
		Message string `json:"message"`
	}{
		Message: "asset deleted successfully",
	})
}
