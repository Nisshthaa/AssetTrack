package services

import (
	"AssetTrack/database"
	"AssetTrack/models"
	"AssetTrack/repository"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func CreateAsset(body models.AssetRequest) (int, string, error) {

	v := validator.New()
	if err := v.Struct(body); err != nil {
		return http.StatusBadRequest, "input validation failed", err
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
		return http.StatusInternalServerError, "transaction failed", txErr
	}

	return http.StatusCreated, "asset created successfully", nil
}

func GetAssets() ([]models.AssetDetails, int, string, error) {

	assets, err := repository.GetAssets()

	if err != nil {
		return nil, http.StatusInternalServerError, "failed to get assets", err
	}

	return assets, http.StatusOK, "assets fetched successfully", nil
}

func GetAssetByID(assetID string) (models.AssetDetails, int, string, error) {

	var asset models.AssetDetails

	txErr := database.Tx(func(tx *sqlx.Tx) error {

		var err error

		asset, err = repository.GetAssetByID(assetID)
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
		return models.AssetDetails{}, http.StatusInternalServerError, "transaction failed", txErr
	}

	return asset, http.StatusOK, "asset fetched successfully", nil
}

func UpdateAsset(assetID string, body models.UpdateAssetRequest) (int, string, error) {

	v := validator.New()
	if err := v.Struct(body); err != nil {
		return http.StatusBadRequest, "input validation failed", err
	}

	txErr := database.Tx(func(tx *sqlx.Tx) error {

		updateErr := repository.UpdateAsset(tx, assetID, body)
		if updateErr != nil {
			return fmt.Errorf("failed to update asset: %w", updateErr)
		}

		switch body.AssetType {

		case "laptop":
			if err := repository.UpdateLaptopSpecs(tx, assetID, body.Laptop); err != nil {
				return fmt.Errorf("failed to update laptop specs: %w", err)
			}

		case "keyboard":
			if err := repository.UpdateKeyboardSpecs(tx, assetID, body.Keyboard); err != nil {
				return fmt.Errorf("failed to update keyboard specs: %w", err)
			}

		case "mouse":
			if err := repository.UpdateMouseSpecs(tx, assetID, body.Mouse); err != nil {
				return fmt.Errorf("failed to update mouse specs: %w", err)
			}

		case "mobile":
			if err := repository.UpdateMobileSpecs(tx, assetID, body.Mobile); err != nil {
				return fmt.Errorf("failed to update mobile specs: %w", err)
			}

		default:
			return fmt.Errorf("unsupported asset type: %s", body.AssetType)
		}

		return nil
	})

	if txErr != nil {

		return http.StatusInternalServerError, "transaction failed", txErr
	}

	return http.StatusOK, "asset updated successfully", nil
}

func DeleteAsset(assetID string) (error, int, string) {

	txErr := database.Tx(func(tx *sqlx.Tx) error {

		if err := repository.ArchiveAsset(tx, assetID); err != nil {
			return fmt.Errorf("failed to archive asset: %w", err)
		}

		if err := repository.UpdateAssetStatus(tx, assetID, "damaged"); err != nil {
			return fmt.Errorf("failed to archive asset: %w", err)
		}

		if err := repository.ArchiveAssetAssignment(tx, assetID); err != nil {
			return fmt.Errorf("failed to archive asset assignments: %w", err)
		}

		return nil
	})

	if txErr != nil {
		return txErr, http.StatusInternalServerError, "failed to delete asset"
	}

	return nil, http.StatusOK, "asset deleted successfully"
}

func AssetSentToRepair(assetID string) (error, int, string) {

	userID, err := repository.GetAssignedUser(assetID)
	if err != nil {
		return err, http.StatusInternalServerError, "failed to get assigned user"
	}

	txErr := database.Tx(func(tx *sqlx.Tx) error {

		if err := repository.ReturnUserAsset(tx, assetID, userID); err != nil {
			return fmt.Errorf("failed to return asset: %w", err)
		}

		if err := repository.AssetSentToRepair(tx, assetID); err != nil {
			return fmt.Errorf("failed to create repair record: %w", err)
		}

		if err := repository.UpdateAssetStatus(tx, assetID, "under_repair"); err != nil {
			return fmt.Errorf("failed to update asset status: %w", err)
		}

		return nil
	})

	if txErr != nil {
		return txErr, http.StatusInternalServerError, "failed to send asset for repair"
	}

	return nil, http.StatusOK, "assent sent to repair"
}

func AssetRepairCompleted(assetID string) (error, int, string) {

	txErr := database.Tx(func(tx *sqlx.Tx) error {

		if err := repository.AssetRepairCompleted(tx, assetID); err != nil {
			return fmt.Errorf("failed to update repair status: %w", err)
		}

		if err := repository.UpdateAssetStatus(tx, assetID, "available"); err != nil {
			return fmt.Errorf("failed to update asset status: %w", err)
		}

		return nil
	})

	if txErr != nil {
		return txErr, http.StatusInternalServerError, "failed to complete asset repair"
	}

	return nil, http.StatusOK, "asset repair status updated"
}

func AdminDashboard() (models.Dashboard, int, string, error) {

	dashboard, err := repository.AdminDashboard()

	if err != nil {
		return models.Dashboard{}, http.StatusInternalServerError, "failed to fetch data", err
	}

	return dashboard, http.StatusOK, "dashboard data fetched successfully", err
}
