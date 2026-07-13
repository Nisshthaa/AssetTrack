package repository

import (
	"AssetTrack/database"
	"AssetTrack/models"
	"AssetTrack/utils"
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
	if err := database.DB.QueryRowx(SQL, Name, Email, Password, PhoneNo, Role, RoleType).Scan(&userID); err != nil {
		return "", err
	}

	return userID, nil
}

func GetUserIDByPassword(email, password string) (string, error) {
	SQL := `SELECT user_id,
       			   password
			  FROM users 
			  WHERE email = TRIM($1)
			    AND archived_at IS NULL`

	var user models.LoginData
	if err := database.DB.Get(&user, SQL, email); err != nil {
		return "", err
	}

	if passwordErr := utils.CheckPassword(password, user.PasswordHash); passwordErr != nil {
		return "", passwordErr
	}
	return user.UserID, nil
}

func GetUser(userID string) (*models.User, error) {
	var user models.User
	SQL := `SELECT user_id, name, email , phone_no, role, type
              FROM users 
              WHERE user_id = $1
                AND archived_at IS NULL`

	getErr := database.DB.Get(&user, SQL, userID)
	return &user, getErr
}

func DeleteUser(userID string) error {
	SQL := `UPDATE users
			  SET archived_at = NOW()
			  WHERE user_id = $1
			    AND archived_at IS NULL`

	_, delErr := database.DB.Exec(SQL, userID)
	return delErr
}
