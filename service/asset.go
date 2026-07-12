package services

import (
	"AssetTrack/database"
	"AssetTrack/models"
	"AssetTrack/repository"
	"AssetTrack/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func CreateAsset(body models.AssetRequest) *models.ServiceResponse {

	v := validator.New()
	if err := v.Struct(body); err != nil {
		return utils.ServiceError(err, http.StatusBadRequest, "input validation failed")
	}

	if body.WarrantyEnd.Before(body.WarrantyStart) {
		return utils.ServiceError(fmt.Errorf("invalid warranty range"), http.StatusBadRequest, "invalid warranty range")
	}

	txErr := database.Tx(func(tx *sqlx.Tx) error {

		assetID, err := repository.CreateAsset(tx, body)
		if err != nil {
			return fmt.Errorf("failed to create asset: %w", err)
		}

		switch body.AssetType {

		case "laptop":
			if err := repository.CreateLaptopSpecs(tx, assetID, body.Laptop); err != nil {
				return fmt.Errorf("failed to create laptop specs: %w", err)
			}

		case "keyboard":
			if err := repository.CreateKeyboardSpecs(tx, assetID, body.Keyboard); err != nil {
				return fmt.Errorf("failed to create keyboard specs: %w", err)
			}

		case "mouse":
			if err := repository.CreateMouseSpecs(tx, assetID, body.Mouse); err != nil {
				return fmt.Errorf("failed to create mouse specs: %w", err)
			}

		case "mobile":
			if err := repository.CreateMobileSpecs(tx, assetID, body.Mobile); err != nil {
				return fmt.Errorf("failed to create mobile specs: %w", err)
			}

		default:
			return fmt.Errorf("unsupported asset type: %s", body.AssetType)
		}

		return nil
	})

	if txErr != nil {
		return utils.ServiceError(txErr, http.StatusInternalServerError, "failed to create asset")
	}

	return utils.ServiceSuccess(nil, http.StatusCreated)
}
