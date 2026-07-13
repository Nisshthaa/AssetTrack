package services

import (
	"AssetTrack/repository"

	"AssetTrack/models"
	"AssetTrack/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func RegisterUser(body models.RegisterUser) (string, int, error) {

	v := validator.New()
	if err := v.Struct(body); err != nil {
		return "", http.StatusBadRequest, err
	}

	exists, existErr := repository.IsUserExists(body.Email)
	if existErr != nil {
		return "", http.StatusInternalServerError, existErr
	}

	if exists {
		return "", http.StatusBadRequest, nil
	}

	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	userID, userErr := repository.CreateUser(
		body.Name,
		body.Email,
		hashedPassword,
		body.PhoneNumber,
		body.Role,
		body.RoleType,
	)
	if userErr != nil {
		return "", http.StatusInternalServerError, userErr
	}

	token, tokenErr := utils.GenerateJWT(userID, body.Role)
	if tokenErr != nil {
		return "", http.StatusInternalServerError, tokenErr
	}

	return token, http.StatusCreated, nil
}

func LoginUser(body models.LoginUser) (string, int, error) {
	v := validator.New()
	if err := v.Struct(body); err != nil {
		return "", http.StatusBadRequest, err
	}

	user, userErr := repository.GetUserByPassword(body.Email, body.Password)
	if userErr != nil {
		return "", http.StatusUnauthorized, userErr
	}

	token, tokenErr := utils.GenerateJWT(user.UserID, user.Role)
	if tokenErr != nil {
		return "", http.StatusInternalServerError, tokenErr
	}

	return token, http.StatusCreated, nil
}

func GetUser(userID string) (*models.User, int, error) {
	user, err := repository.GetUser(userID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}

func DeleteUser(userID string) (int, error) {
	err := repository.DeleteUser(userID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
