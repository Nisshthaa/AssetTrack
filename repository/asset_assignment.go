package repository

import (
	"AssetTrack/database"
	"AssetTrack/models"

	"github.com/jmoiron/sqlx"
)

func AssignAsset(tx *sqlx.Tx, assetID, userID string) error {

	SQL := `INSERT INTO asset_assignments (asset_id,assigned_to)
		VALUES ($1,$2);`

	_, err := tx.Exec(SQL, assetID, userID)
	return err
}

func UpdateAssetStatus(tx *sqlx.Tx, assetID, status string) error {

	SQL := `UPDATE assets
		SET status = $1,
			updated_at = CURRENT_TIMESTAMP
		WHERE asset_id = $2
			AND archived_at IS NULL;`

	_, err := tx.Exec(SQL, status, assetID)
	return err
}

func ReturnUserAsset(tx *sqlx.Tx, assetID, userID string) error {

	SQL := `
		UPDATE asset_assignments
		SET
			returned_at = CURRENT_TIMESTAMP,
			updated_at = CURRENT_TIMESTAMP
		WHERE
			asset_id = $1
			AND assigned_to = $2
			AND returned_at IS NULL
			AND archived_at IS NULL;
	`

	_, err := tx.Exec(SQL, assetID, userID)
	if err != nil {
		return err
	}

	return nil
}

func GetAssetHistory(assetID string) ([]models.AssetHistory, error) {

	var history []models.AssetHistory

	SQL := `SELECT a.asset_id,a.asset_type,a.brand,a.model,u.user_id,u.name,aa.assigned_on,aa.returned_at
			FROM assets a 
			JOIN asset_assignments aa ON a.asset_id=aa.asset_id
    		JOIN users u ON u.user_id=aa.assigned_to 
			WHERE aa.asset_id=$1
			AND aa.archived_at IS NULL
			AND a.archived_at IS NULL
			AND u.archived_at IS NULL
			ORDER BY aa.assigned_on DESC;`

	err := database.DB.Select(&history, SQL, assetID)
	if err != nil {
		return nil, err
	}

	return history, nil
}
