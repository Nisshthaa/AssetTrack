package services

import (
	"AssetTrack/database"
	"AssetTrack/models"
	"AssetTrack/repository"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

func AssignAsset(body models.AssignAssetRequest) (error, int, string) {

	v := validator.New()
	if err := v.Struct(body); err != nil {
		return err, http.StatusBadRequest, "input validation failed"
	}

	txErr := database.Tx(func(tx *sqlx.Tx) error {

		if assignErr := repository.AssignAsset(tx, body.AssetID, body.UserID); assignErr != nil {
			return fmt.Errorf("failed to assign asset: %w", assignErr)
		}

		if updateErr := repository.UpdateAssetStatus(tx, body.AssetID, "assigned"); updateErr != nil {
			return fmt.Errorf("failed to update asset status: %w", updateErr)
		}

		return nil
	})

	if txErr != nil {
		return txErr, http.StatusInternalServerError, "failed to assign asset"
	}

	return nil, http.StatusCreated, "asset assigned successfully"
}

func ReturnAsset(body models.AssignAssetRequest) (error, int, string) {

	v := validator.New()
	if err := v.Struct(body); err != nil {
		return err, http.StatusBadRequest, "input validation failed"
	}

	txErr := database.Tx(func(tx *sqlx.Tx) error {

		if err := repository.ReturnUserAsset(tx, body.AssetID, body.UserID); err != nil {
			return fmt.Errorf("failed to return asset: %w", err)
		}

		if err := repository.UpdateAssetStatus(tx, body.AssetID, "available"); err != nil {
			return fmt.Errorf("failed to update asset status: %w", err)
		}

		return nil
	})

	if txErr != nil {

		if errors.Is(txErr, sql.ErrNoRows) {
			return txErr, http.StatusNotFound, "asset assignment not found"
		}

		return txErr, http.StatusInternalServerError, "failed to return asset"
	}

	return nil, http.StatusOK, "asset returned successfully"
}
