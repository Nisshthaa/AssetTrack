package repository

import (
	"AssetTrack/database"
	"AssetTrack/models"
	"AssetTrack/utils"

	"github.com/jmoiron/sqlx"
)

func IsUserExists(email string) (bool, error) {
	SQL := `SELECT count(user_id) > 0 as is_exist
			  FROM users
			  WHERE email = TRIM($1)
			  AND archived_at IS NULL`

	var exist bool
	err := database.DB.Get(&exist, SQL, email)
	return exist, err
}

func CreateUser(Name, Email, Password, PhoneNo, Role, RoleType string) (string, error) {

	SQL := `INSERT INTO users(name, email, password, phone_no, role, type)
			VALUES ($1, TRIM(LOWER($2)), $3, $4, $5, $6)
			RETURNING user_id`
	var userID string
	crtErr := database.DB.QueryRowx(SQL, Name, Email, Password, PhoneNo, Role, RoleType).Scan(&userID)

	return userID, crtErr
}

func GetUserByPassword(body models.LoginRequest) (models.LoginData, error) {
	SQL := `SELECT user_id, password, role
			  FROM users 
			  WHERE email = TRIM($1)
			  AND archived_at IS NULL`

	var user models.LoginData
	if err := database.DB.Get(&user, SQL, body.Email); err != nil {
		return models.LoginData{}, err
	}

	if passwordErr := utils.CheckPassword(body.Password, user.PasswordHash); passwordErr != nil {
		return models.LoginData{}, passwordErr
	}
	return user, nil
}

func GetUser(userID string) (models.User, error) {
	var user models.User

	SQL := `SELECT user_id, name, email , phone_no, role, type,created_at
            FROM users 
            WHERE user_id = $1
            AND archived_at IS NULL`

	getErr := database.DB.Get(&user, SQL, userID)
	return user, getErr
}

func GetUserAssets(userID string) ([]models.AssetDetails, error) {
	var assets []models.AssetDetails

	SQL := `SELECT a.asset_id,a.serial_number,a.brand,a.model,a.asset_type,a.status,a.owner_type,a.warranty_start,a.warranty_end 
			FROM asset_assignments aa 
			JOIN assets a 
			ON aa.asset_id = a.asset_id
			WHERE aa.assigned_to = $1 
			AND aa.returned_at IS NULL 
			AND aa.archived_at IS NULL 
			AND a.archived_at IS NULL; `

	err := database.DB.Select(&assets, SQL, userID)
	return assets, err
}

func GetUserAssetByID(tx *sqlx.Tx, userID, assetID string) (models.AssetDetails, error) {
	var asset models.AssetDetails

	SQL := `SELECT a.asset_id,a.serial_number,a.brand,a.model,a.asset_type,a.status,a.owner_type,a.warranty_start,a.warranty_end
          FROM assets a 
          JOIN asset_assignments aa ON aa.asset_id=a.asset_id
          WHERE aa.assigned_to=$1 
          AND aa.asset_id=$2
          AND aa.returned_at IS NULL
          AND aa.archived_at IS NULL
          AND a.archived_at IS NULL`

	err := tx.Get(&asset, SQL, userID, assetID)
	return asset, err
}
