package repository

import (
	"AssetTrack/database"

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

func GetAssignedUser(assetID string) (string, error) {
	var userID string

	SQL := `SELECT assigned_to
		FROM asset_assignments
		WHERE
			asset_id = $1
			AND returned_at IS NULL
			AND archived_at IS NULL;`
	err := database.DB.Get(&userID, SQL, assetID)
	return userID, err
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

func ReturnUserAssets(tx *sqlx.Tx, userID string) (string, error) {

	SQL := `UPDATE asset_assignments
		  SET returned_at=NOW()
		  WHERE assigned_to=$1
		  AND returned_at IS NULL
		  RETURNING asset_id`

	var assetID string

	err := tx.QueryRowx(SQL, userID).Scan(&assetID)
	if err != nil {
		return "", err
	}
	return assetID, nil
}
