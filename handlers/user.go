package handlers

import (
	"AssetTrack/middlewares"
	"AssetTrack/models"
	services "AssetTrack/service"
	"AssetTrack/utils"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var body models.RegisterUser

	if parseErr := utils.ParseBody(r, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "invalid request body")
		return
	}

	token, statusCode, err := services.RegisterUser(body)

	if err != nil {
		utils.RespondError(w, statusCode, err, "user not registered")
		return
	}

	utils.RespondJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}{
		Message: "User registered successfully",
		Token:   token,
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var body models.LoginRequest

	if parseErr := utils.ParseBody(r, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	token, statusCode, err := services.LoginUser(body)

	if err != nil {
		utils.RespondError(w, statusCode, err, err.Error())
		return
	}

	if token == "" {
		utils.RespondError(w, statusCode, nil, "user already exists")
		return
	}

	utils.RespondJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}{
		Message: "User logged in successfully",
		Token:   token,
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userCtx := middlewares.GetUserContext(r)

	user, statusCode, message, err := services.GetUser(userCtx.UserID)
	if err != nil {
		utils.RespondError(w, statusCode, err, message)
		return
	}

	utils.RespondJSON(w, statusCode, user)
}
func GetUserAssets(w http.ResponseWriter, r *http.Request) {
	userCtx := middlewares.GetUserContext(r)

	assets, statusCode, message, err := services.GetUserAssets(userCtx.UserID)
	if err != nil {
		utils.RespondError(w, statusCode, err, message)
		return
	}
	utils.RespondJSON(w, statusCode, assets)
}

func GetUserAssetByID(w http.ResponseWriter, r *http.Request) {
	userCtx := middlewares.GetUserContext(r)
	assetID := r.PathValue("assetID")

	assets, statusCode, message, err := services.GetUserAssetByID(userCtx.UserID, assetID)
	if err != nil {
		utils.RespondError(w, statusCode, err, message)
		return
	}
	utils.RespondJSON(w, statusCode, assets)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {

	utils.RespondJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{"logout successful"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userCtx := middlewares.GetUserContext(r)

	status, message, err := services.DeleteUser(userCtx.UserID)
	if err != nil {
		utils.RespondError(w, status, err, message)
		return
	}

	utils.RespondJSON(w, status, struct {
		Message string `json:"message"`
	}{
		Message: "account deleted successfully",
	})
}
