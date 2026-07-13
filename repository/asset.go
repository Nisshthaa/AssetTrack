package repository

import (
	"AssetTrack/database"
	"AssetTrack/models"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func CreateAsset(tx *sqlx.Tx, body models.AssetRequest) (string, error) {
	var assetID string

	SQL := `INSERT INTO assets (serial_number,brand,model,asset_type,status,owner_type,warranty_start,warranty_end)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING asset_id;`

	err := tx.QueryRow(SQL, body.SerialNumber, body.Brand, body.Model, body.AssetType, body.Status, body.OwnerType, body.WarrantyStart, body.WarrantyEnd).
		Scan(&assetID)

	return assetID, err
}

func CreateLaptopSpecs(tx *sqlx.Tx, assetID string, body models.LaptopSpecsRequest) error {

	SQL := `INSERT INTO laptop (asset_id,processor,ram,storage,operating_system,charger)
			VALUES ($1,$2,$3,$4,$5,$6)`

	_, err := tx.Exec(SQL, assetID, body.Processor, body.Ram, body.Storage, body.OperatingSystem, body.Charger)

	return err
}

func CreateKeyboardSpecs(tx *sqlx.Tx, assetID string, body models.KeyboardSpecsRequest) error {

	SQL := `INSERT INTO keyboard (asset_id,layout,connection_type)
			VALUES ($1,$2,$3)`

	_, err := tx.Exec(SQL, assetID, body.Layout, body.ConnectionType)

	return err
}

func CreateMouseSpecs(tx *sqlx.Tx, assetID string, body models.MouseSpecsRequest) error {

	SQL := `INSERT INTO mouse (asset_id,dpi,connection_type)
			VALUES ($1,$2,$3)`

	_, err := tx.Exec(SQL, assetID, body.Dpi, body.ConnectionType)

	return err
}

func CreateMobileSpecs(tx *sqlx.Tx, assetID string, body models.MobileSpecsRequest) error {

	SQL := `INSERT INTO mobile (asset_id,ram,storage,operating_system,charger)
			VALUES ($1,$2,$3,$4,$5)`

	_, err := tx.Exec(SQL, assetID, body.Ram, body.Storage, body.OperatingSystem, body.Charger)

	return err
}

func GetAssets() ([]models.AssetDetails, error) {
	var assets []models.AssetDetails

	SQL := `
		SELECT
			a.asset_id,
			a.serial_number,
			a.brand,
			a.model,
			a.asset_type,
			a.status,
			a.owner_type,
			COALESCE(u.name, '') AS assigned_to,
			a.warranty_start,
			a.warranty_end
		FROM assets a
		LEFT JOIN asset_assignments aa
			ON aa.asset_id = a.asset_id
			AND aa.returned_at IS NULL
			AND aa.archived_at IS NULL
		LEFT JOIN users u
			ON u.user_id = aa.assigned_to
			AND u.archived_at IS NULL
		WHERE a.archived_at IS NULL
		ORDER BY a.created_at DESC;
	`

	err := database.DB.Select(&assets, SQL)

	return assets, err
}

func GetAssetByID(assetID string) (models.AssetDetails, error) {

	var asset models.AssetDetails

	SQL := `SELECT a.asset_id,a.serial_number,a.brand,a.model,a.asset_type,a.status,a.owner_type,COALESCE(u.name, '') AS assigned_to,a.warranty_start,a.warranty_end
			FROM assets a
			LEFT JOIN asset_assignments aa
			ON aa.asset_id = a.asset_id
			AND aa.returned_at IS NULL
			AND aa.archived_at IS NULL
			LEFT JOIN users u
			ON u.user_id = aa.assigned_to
			AND u.archived_at IS NULL
			WHERE a.asset_id = $1
			AND a.archived_at IS NULL`

	err := database.DB.Get(&asset, SQL, assetID)
	if err != nil {
		return models.AssetDetails{}, err
	}

	return asset, nil
}

func UpdateAsset(tx *sqlx.Tx, assetID string, body models.UpdateAssetRequest) (string, error) {
	var assetType string

	SQL := `UPDATE assets
		SET serial_number = $1,brand = $2,model = $3,status = $4,owner_type = $5,warranty_start = $6,warranty_end = $7,updated_at = CURRENT_TIMESTAMP
		WHERE asset_id = $8
			AND archived_at IS NULL
			RETURNING asset_type`

	if err := tx.QueryRowx(SQL, body.SerialNumber, body.Brand, body.Model, body.Status, body.OwnerType, body.WarrantyStart, body.WarrantyEnd, assetID).Scan(&assetType); err != nil {
		return "", err
	}
	return assetType, nil
}

func UpdateLaptopSpecs(tx *sqlx.Tx, assetID string, body models.LaptopSpecsRequest) error {

	SQL := `UPDATE laptop
		SET processor = $1,ram = $2,storage = $3,operating_system = $4,charger = $5
		WHERE asset_id = $6;`

	_, err := tx.Exec(SQL, body.Processor, body.Ram, body.Storage, body.OperatingSystem, body.Charger, assetID)

	return err
}

func UpdateKeyboardSpecs(tx *sqlx.Tx, assetID string, body models.KeyboardSpecsRequest) error {

	SQL := `UPDATE keyboard
		SET layout = $1,connection_type = $2
		WHERE asset_id = $3;`

	_, err := tx.Exec(
		SQL, body.Layout, body.ConnectionType, assetID)

	return err
}

func UpdateMouseSpecs(tx *sqlx.Tx, assetID string, body models.MouseSpecsRequest) error {

	SQL := `UPDATE mouse
		SET dpi = $1,connection_type = $2
		WHERE asset_id = $3;`

	_, err := tx.Exec(SQL, body.Dpi, body.ConnectionType, assetID)

	return err
}

func UpdateMobileSpecs(tx *sqlx.Tx, assetID string, body models.MobileSpecsRequest) error {

	SQL := `UPDATE mobile
		SET operating_system = $1,ram = $2,storage = $3,charger = $4
		WHERE asset_id = $5;`

	_, err := tx.Exec(SQL, body.OperatingSystem, body.Ram, body.Storage, body.Charger, assetID)

	return err
}

func ArchiveAsset(tx *sqlx.Tx, assetID string) error {

	SQL := `UPDATE assets
		SET archived_at = CURRENT_TIMESTAMP
		WHERE asset_id = $1
			AND archived_at IS NULL;`

	result, err := tx.Exec(SQL, assetID)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func ArchiveAssetAssignment(tx *sqlx.Tx, assetID string) error {

	SQL := `UPDATE asset_assignments
		SET archived_at = CURRENT_TIMESTAMP
		WHERE asset_id = $1
			AND archived_at IS NULL;
	`

	_, err := tx.Exec(SQL, assetID)
	return err
}
