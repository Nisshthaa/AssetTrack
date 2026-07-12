package repository

import (
	"AssetTrack/models"

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
