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

func CreateLaptopSpecs(tx *sqlx.Tx, assetID string, body models.LaptopSpecs) error {

	SQL := `INSERT INTO laptop (asset_id,processor,ram,storage,operating_system,charger)
			VALUES ($1,$2,$3,$4,$5,$6)`

	_, err := tx.Exec(SQL, assetID, body.Processor, body.Ram, body.Storage, body.OperatingSystem, body.Charger)

	return err
}

func CreateKeyboardSpecs(tx *sqlx.Tx, assetID string, body models.KeyboardSpecs) error {

	SQL := `INSERT INTO keyboard (asset_id,layout,connection_type)
			VALUES ($1,$2,$3)`

	_, err := tx.Exec(SQL, assetID, body.Layout, body.ConnectionType)

	return err
}

func CreateMouseSpecs(tx *sqlx.Tx, assetID string, body models.MouseSpecs) error {

	SQL := `INSERT INTO mouse (asset_id,dpi,connection_type)
			VALUES ($1,$2,$3)`

	_, err := tx.Exec(SQL, assetID, body.Dpi, body.ConnectionType)

	return err
}

func CreateMobileSpecs(tx *sqlx.Tx, assetID string, body models.MobileSpecs) error {

	SQL := `INSERT INTO mobile (asset_id,ram,storage,operating_system,charger)
			VALUES ($1,$2,$3,$4,$5)`

	_, err := tx.Exec(SQL, assetID, body.Ram, body.Storage, body.OperatingSystem, body.Charger)

	return err
}

func GetAssets() ([]models.AssetDetails, error) {
	var assets []models.AssetDetails

	SQL := `SELECT asset_id,serial_number,brand,model,asset_type,status,owner_type,warranty_start,warranty_end
		    FROM assets 
			WHERE archived_at IS NULL`

	err := database.DB.Select(&assets, SQL)
	return assets, err
}

func GetAssetByID(assetID string) (models.AssetDetails, error) {

	var asset models.AssetDetails

	SQL := `SELECT asset_id,serial_number,brand,model,asset_type,status,owner_type,warranty_start,warranty_end
			FROM assets 
			WHERE asset_id = $1
			AND archived_at IS NULL`

	err := database.DB.Get(&asset, SQL, assetID)
	if err != nil {
		return models.AssetDetails{}, err
	}

	return asset, nil
}

func GetLaptopSpecifications(tx *sqlx.Tx, assetID string) (models.LaptopSpecs, error) {

	var laptop models.LaptopSpecs

	SQL := `SELECT processor,ram,storage ,operating_system ,Charger
			FROM laptop 
			WHERE asset_id = $1`

	err := tx.Get(&laptop, SQL, assetID)
	if err != nil {
		return models.LaptopSpecs{}, err
	}

	return laptop, nil
}

func GetKeyboardSpecifications(tx *sqlx.Tx, assetID string) (models.KeyboardSpecs, error) {

	var keyboard models.KeyboardSpecs

	SQL := `SELECT layout,connection_type
			FROM keyboard 
			WHERE asset_id = $1`

	err := tx.Get(&keyboard, SQL, assetID)
	if err != nil {
		return models.KeyboardSpecs{}, err
	}

	return keyboard, nil
}

func GetMouseSpecifications(tx *sqlx.Tx, assetID string) (models.MouseSpecs, error) {

	var mouse models.MouseSpecs

	SQL := `SELECT dpi,connection_type
			FROM  mouse 
			WHERE asset_id = $1`

	err := tx.Get(&mouse, SQL, assetID)
	if err != nil {
		return models.MouseSpecs{}, err
	}

	return mouse, nil
}

func GetMobileSpecifications(tx *sqlx.Tx, assetID string) (models.MobileSpecs, error) {

	var mobile models.MobileSpecs

	SQL := `SELECT ram,storage ,operating_system ,Charger
			FROM mobile 
			WHERE asset_id = $1`

	err := tx.Get(&mobile, SQL, assetID)
	if err != nil {
		return models.MobileSpecs{}, err
	}

	return mobile, nil
}
func UpdateAsset(tx *sqlx.Tx, assetID string, body models.UpdateAssetRequest) error {

	SQL := `UPDATE assets
			SET serial_number = $1,brand = $2,model = $3,status = $4,owner_type = $5,warranty_start = $6,warranty_end = $7,updated_at = CURRENT_TIMESTAMP
			WHERE asset_id = $8
			AND archived_at IS NULL`

	_, err := tx.Exec(SQL, body.SerialNumber, body.Brand, body.Model, body.Status, body.OwnerType, body.WarrantyStart, body.WarrantyEnd, assetID)

	return err
}

func UpdateLaptopSpecs(tx *sqlx.Tx, assetID string, body models.LaptopSpecs) error {

	SQL := `UPDATE laptop
		SET processor =$1,ram=$2,storage =$3,operating_system = $4,charger = $5
		WHERE asset_id = $6;`

	_, err := tx.Exec(SQL, body.Processor, body.Ram, body.Storage, body.OperatingSystem, body.Charger, assetID)

	return err
}

func UpdateKeyboardSpecs(tx *sqlx.Tx, assetID string, body models.KeyboardSpecs) error {

	SQL := `UPDATE keyboard
		SET layout = $1,connection_type=$2
		WHERE asset_id = $3;`

	_, err := tx.Exec(
		SQL, body.Layout, body.ConnectionType, assetID)

	return err
}

func UpdateMouseSpecs(tx *sqlx.Tx, assetID string, body models.MouseSpecs) error {

	SQL := `UPDATE mouse
		SET dpi =$1,connection_type = $2
		WHERE asset_id = $3;`

	_, err := tx.Exec(SQL, body.Dpi, body.ConnectionType, assetID)

	return err
}

func UpdateMobileSpecs(tx *sqlx.Tx, assetID string, body models.MobileSpecs) error {

	SQL := `UPDATE mobile
		SET operating_system =$1,ram = $2,storage = $3,charger = $4
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

func AssetSentToRepair(tx *sqlx.Tx, assetID string) error {
	SQL := `INSERT INTO asset_repairs(asset_id)
			VALUES ($1)`

	_, err := tx.Exec(SQL, assetID)

	return err
}

func AssetRepairCompleted(tx *sqlx.Tx, assetID string) error {
	SQL := `UPDATE asset_repairs
			SET repair_completed_on=NOW(),updated_at=NOW()
			WHERE asset_id=$1
			AND archived_at is NULL`

	_, err := tx.Exec(SQL, assetID)

	return err
}

func AdminDashboard() (models.Dashboard, error) {

	var dashboard models.Dashboard

	SQL := `SELECT COUNT(*) as totalAssets,
       		COUNT(CASE when status='assigned' THEN 1 END) AS totalAssignedAssets,
       		COUNT(CASE when status='available' THEN 1 END) AS totalAvailableAssets,
       		COUNT(CASE when status='needs_repair' THEN 1 END) AS needsRepairAssets,
       		COUNT(CASE when status='under_repair' THEN 1 END) AS underRepairAssets,
       		COUNT(CASE when status='damaged' THEN 1 END) AS damagedAssets
       		FROM assets
			WHERE archived_at IS NULL`

	if err := database.DB.Get(&dashboard, SQL); err != nil {
		return models.Dashboard{}, err
	}
	return dashboard, nil
}
