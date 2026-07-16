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

func GetLaptopSpecifications(tx *sqlx.Tx, assetID string) (models.LaptopSpecsRequest, error) {

	var laptop models.LaptopSpecsRequest

	SQL := `SELECT processor,ram,storage ,operating_system ,Charger
			FROM laptop 
			WHERE asset_id = $1`

	err := tx.Get(&laptop, SQL, assetID)
	if err != nil {
		return models.LaptopSpecsRequest{}, err
	}

	return laptop, nil
}

func GetKeyboardSpecifications(tx *sqlx.Tx, assetID string) (models.KeyboardSpecsRequest, error) {

	var keyboard models.KeyboardSpecsRequest

	SQL := `SELECT layout,connection_type
			FROM keyboard 
			WHERE asset_id = $1`

	err := tx.Get(&keyboard, SQL, assetID)
	if err != nil {
		return models.KeyboardSpecsRequest{}, err
	}

	return keyboard, nil
}

func GetMouseSpecifications(tx *sqlx.Tx, assetID string) (models.MouseSpecsRequest, error) {

	var mouse models.MouseSpecsRequest

	SQL := `SELECT dpi,connection_type
			FROM  mouse 
			WHERE asset_id = $1`

	err := tx.Get(&mouse, SQL, assetID)
	if err != nil {
		return models.MouseSpecsRequest{}, err
	}

	return mouse, nil
}

func GetMobileSpecifications(tx *sqlx.Tx, assetID string) (models.MobileSpecsRequest, error) {

	var mobile models.MobileSpecsRequest

	SQL := `SELECT ram,storage ,operating_system ,Charger
			FROM mobile 
			WHERE asset_id = $1`

	err := tx.Get(&mobile, SQL, assetID)
	if err != nil {
		return models.MobileSpecsRequest{}, err
	}

	return mobile, nil
}
func UpdateAsset(tx *sqlx.Tx, assetID string, body models.UpdateAssetRequest) (string, error) {
	var assetType string

	SQL := `UPDATE assets
			SET serial_number =COALESCE($1,serial_number),
			    brand = COALESCE($2,brand),
			    model = COALESCE($3,model),
			    status = COALESCE($4,status),
			    owner_type = COALESCE($5,owner_type),
			    warranty_start = COALESCE($6,warranty_start),
			    warranty_end =COALESCE($7,warranty_end),
			    updated_at = CURRENT_TIMESTAMP
			WHERE asset_id = $8
			AND archived_at IS NULL
			RETURNING asset_type`

	if err := tx.QueryRowx(SQL, body.SerialNumber, body.Brand, body.Model, body.Status, body.OwnerType, body.WarrantyStart, body.WarrantyEnd, assetID).
		Scan(&assetType); err != nil {
		return "", err
	}
	return assetType, nil
}

func UpdateLaptopSpecs(tx *sqlx.Tx, assetID string, body models.LaptopSpecsRequest) error {

	SQL := `UPDATE laptop
		SET processor =COALESCE($1,processor),
		    ram=COALESCE($2,ram),
		    storage =COALESCE($3,storage),
		    operating_system = COALESCE($4,operating_system),
		    charger = COALESCE($5,charger)
		WHERE asset_id = $6;`

	_, err := tx.Exec(SQL, body.Processor, body.Ram, body.Storage, body.OperatingSystem, body.Charger, assetID)

	return err
}

func UpdateKeyboardSpecs(tx *sqlx.Tx, assetID string, body models.KeyboardSpecsRequest) error {

	SQL := `UPDATE keyboard
		SET layout = COALESCE($1,layout),
		    connection_type=COALESCE($2,connection_type)
		WHERE asset_id = $3;`

	_, err := tx.Exec(
		SQL, body.Layout, body.ConnectionType, assetID)

	return err
}

func UpdateMouseSpecs(tx *sqlx.Tx, assetID string, body models.MouseSpecsRequest) error {

	SQL := `UPDATE mouse
		SET dpi =COALESCE($1,dpi),
		    connection_type =COALESCE($2,connection_type)
		WHERE asset_id = $3;`

	_, err := tx.Exec(SQL, body.Dpi, body.ConnectionType, assetID)

	return err
}

func UpdateMobileSpecs(tx *sqlx.Tx, assetID string, body models.MobileSpecsRequest) error {

	SQL := `UPDATE mobile
		SET operating_system =COALESCE($1,operating_system),
		    ram =COALESCE($2,ram),
		    storage = COALESCE($3,storage),
		    charger = COALESCE($4,charger)
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
	SQL := `UPDATE asset_repairs
			SET sent_for_repair_on=NOW(),updated_at=NOW()
			WHERE asset_id=$1
			AND archived_at is NULL`

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
