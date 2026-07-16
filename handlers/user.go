package handlers

import (
	"AssetTrack/database"
	"AssetTrack/database/dbHelper"
	"AssetTrack/models"
	"AssetTrack/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var body models.RegisterUser
	if parseErr := utils.ParseBody(r, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	v := validator.New()
	if err := v.Struct(body); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "input validation failed")
		return
	}

	exists, existErr := dbHelper.IsUserExists(body.Email)
	if existErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, existErr, "failed to check user existence")
		return
	}

	if exists {
		utils.RespondError(w, http.StatusBadRequest, nil, "user already exists")
		return
	}

	hashedPassword, hashErr := utils.HashPassword(body.Password)
	if hashErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, hashErr, "failed to secure password")
		return
	}

	userID, saveErr := dbHelper.CreateUser(database.DB, body.Name, body.Email, hashedPassword, body.PhoneNumber, body.Role, body.RoleType)
	if saveErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, saveErr, "failed to create user")
		return
	}

	token, tokenErr := utils.GenerateJWT(userID)
	if tokenErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, tokenErr, "failed to generate token")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}{
		Message: "User registered successfully",
		Token:   token,
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var body models.LoginUser

	if parseErr := utils.ParseBody(r, &body); parseErr != nil {
		utils.RespondError(w, http.StatusBadRequest, parseErr, "failed to parse request body")
		return
	}

	v := validator.New()
	if err := v.Struct(body); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "input validation failed")
		return
	}

	userID, userErr := dbHelper.GetUserIDByPassword(body.Email, body.Password)
	if userErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, userErr, "failed to find user")
		return
	}

	if userID == "" {
		utils.RespondError(w, http.StatusBadRequest, nil, "user not found")
		return
	}

	token, tokenErr := utils.GenerateJWT(userID)
	if tokenErr != nil {
		utils.RespondError(w, http.StatusInternalServerError, tokenErr, "failed to generate token")
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
<<<<<<< Updated upstream
=======

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
>>>>>>> Stashed changes
