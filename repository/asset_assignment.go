package repository

import (
	"database/sql"

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

func ReturnAsset(tx *sqlx.Tx, assetID, userID string) error {

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

	result, err := tx.Exec(SQL, assetID, userID)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
