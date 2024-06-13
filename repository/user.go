package repository

import (
	"database/sql"
	"go-crud/model"
)

func GetUserByUsername(db *sql.DB, username string) (model.User, error) {
	var user model.User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}

func GetUserBySessionToken(db *sql.DB, sessionToken string) (model.User, error) {
	var user model.User
	err := db.QueryRow("SELECT id, username FROM users WHERE session_token = $1", sessionToken).Scan(&user.ID, &user.Username)
	return user, err
}

func UpdateUserSessionToken(db *sql.DB, userID int, sessionToken string) error {
	_, err := db.Exec("UPDATE users SET session_token = $1 WHERE id = $2", sessionToken, userID)
	return err
}

func ClearUserSessionToken(db *sql.DB, sessionToken string) error {
	_, err := db.Exec("UPDATE users SET session_token = NULL WHERE session_token = $1", sessionToken)
	return err
}
