package dbHelper

import (
	"AssetTrack/database"

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

func CreateUser(db sqlx.Ext, Name, Email, Password, PhoneNo, Role, RoleType string) (string, error) {

	SQL := `INSERT INTO users(name, email, password, phone_no, role, type)
			VALUES ($1, TRIM(LOWER($2)), $3, $4, $5, $6)
			RETURNING user_id
`
	var userID string
	if err := db.QueryRowx(SQL, Name, Email, Password, PhoneNo, Role, RoleType).Scan(&userID); err != nil {
		return "", err
	}

	return userID, nil
}
