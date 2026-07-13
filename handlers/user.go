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

	token, status, err := services.RegisterUser(body)

	if err != nil {
		utils.RespondError(w, status, err, err.Error())
		return
	}

	if token == "" {
		utils.RespondError(w, status, nil, "user already exists")
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

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var body models.LoginUser

	if parseErr := utils.ParseBody(r, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	token, status, err := services.LoginUser(body)

	if err != nil {
		utils.RespondError(w, status, err, err.Error())
		return
	}

	if token == "" {
		utils.RespondError(w, status, nil, "user already exists")
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

	user, status, err := services.GetUser(userCtx.UserID)
	if err != nil {
		utils.RespondError(w, status, err, "failed to get user")
		return
	}

	utils.RespondJSON(w, status, user)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {

	utils.RespondJSON(w, http.StatusOK, struct {
		Message string `json:"message"`
	}{"logout successful"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userCtx := middlewares.GetUserContext(r)

	status, err := services.DeleteUser(userCtx.UserID)
	if err != nil {
		utils.RespondError(w, status, err, "failed to delete user account")
		return
	}

	utils.RespondJSON(w, status, struct {
		Message string `json:"message"`
	}{
		Message: "account deleted successfully",
	})
}
