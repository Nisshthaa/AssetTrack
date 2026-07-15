package services

import (
	"AssetTrack/database"
	"AssetTrack/repository"
	"fmt"
	"strings"

	"AssetTrack/models"
	"AssetTrack/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func RegisterUser(body models.RegisterUser) (string, int, error) {

	v := validator.New()
	if err := v.Struct(body); err != nil {
		return "input validation failed", http.StatusBadRequest, err
	}

	exists, existErr := repository.IsUserExists(body.Email)
	if existErr != nil {
		return "failed to check user existence", http.StatusInternalServerError, existErr
	}

	if exists {
		return "user already exists", http.StatusBadRequest, nil
	}

	hashedPassword, hashErr := utils.HashPassword(body.Password)
	if hashErr != nil {
		return "failed to hash password", http.StatusInternalServerError, hashErr
	}

	userID, userErr := repository.CreateUser(body.Name, body.Email, hashedPassword, body.PhoneNumber, body.Role, body.RoleType)
	if userErr != nil {
		return "failed to create user", http.StatusInternalServerError, userErr
	}

	token, tokenErr := utils.GenerateJWT(userID, body.Role)
	if tokenErr != nil {
		return "failed to create token", http.StatusInternalServerError, tokenErr
	}

	return token, http.StatusCreated, nil
}

func LoginUser(body models.LoginRequest) (string, int, error) {

	v := validator.New()
	if err := v.Struct(body); err != nil {
		return "input validation failed", http.StatusBadRequest, err
	}

	user, userErr := repository.GetUserByPassword(body)
	if userErr != nil {
		return "failed to get user", http.StatusUnauthorized, userErr
	}

	token, tokenErr := utils.GenerateJWT(user.UserID, user.Role)
	if tokenErr != nil {
		return "failed to generate token", http.StatusInternalServerError, tokenErr
	}

	return token, http.StatusCreated, nil
}

func GetUser(userID string) (models.User, int, string, error) {

	user, err := repository.GetUser(userID)
	if err != nil {
		return models.User{}, http.StatusInternalServerError, "failed to fetch user", err
	}

	return user, http.StatusOK, "user fetched successfully", nil
}

func GetUserAssets(userID string) ([]models.AssetDetails, int, string, error) {

	assets, err := repository.GetUserAssets(userID)
	if err != nil {
		return nil, http.StatusInternalServerError, "failed to fetch user assets", err
	}

	return assets, http.StatusOK, "user assets fetched successfully", err
}

func GetUserAssetByID(userID, assetID string) (models.AssetDetails, int, string, error) {

	var asset models.AssetDetails

	txErr := database.Tx(func(tx *sqlx.Tx) error {
		var err error

		asset, err = repository.GetUserAssetByID(tx, userID, assetID)
		if err != nil {
			return fmt.Errorf("failed to get asset details: %w", err)
		}

		switch asset.AssetType {

		case "laptop":
			laptop, err := repository.GetLaptopSpecifications(tx, assetID)
			if err != nil {
				return fmt.Errorf("failed to get laptop specs: %w", err)
			}
			asset.Laptop = &laptop

		case "keyboard":
			keyboard, err := repository.GetKeyboardSpecifications(tx, assetID)
			if err != nil {
				return fmt.Errorf("failed to get keyboard specs: %w", err)
			}
			asset.Keyboard = &keyboard

		case "mouse":
			mouse, err := repository.GetMouseSpecifications(tx, assetID)
			if err != nil {
				return fmt.Errorf("failed to get mouse specs: %w", err)
			}
			asset.Mouse = &mouse

		case "mobile":
			mobile, err := repository.GetMobileSpecifications(tx, assetID)
			if err != nil {
				return fmt.Errorf("failed to get mobile specs: %w", err)
			}
			asset.Mobile = &mobile

		default:
			return fmt.Errorf("unsupported asset type: %s", asset.AssetType)
		}

		return nil
	})

	if txErr != nil {
		if strings.Contains(txErr.Error(), "not found") {
			return models.AssetDetails{}, http.StatusNotFound, "asset not found", txErr
		}

		return models.AssetDetails{}, http.StatusInternalServerError, "failed to fetch asset", txErr
	}

	return asset, http.StatusOK, "asset fetched successfully", nil
}

func DeleteUser(userID string) (int, string, error) {
	err := repository.DeleteUser(userID)
	if err != nil {
		return http.StatusInternalServerError, "failed to delete user", err
	}

	return http.StatusOK, "user deleted successfully", nil
}
