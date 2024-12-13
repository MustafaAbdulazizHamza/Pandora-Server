package APIs

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/MustafaAbdulazizHamza/Pandora/Structures"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	errInternalServerError = errors.New("internal server error")
	errNotFound            = errors.New("the requested resource was not found")
)

func execSQL(db *sql.DB, statement string, args ...interface{}) error {
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(args...)
	return err

}
func addSecret(db *sql.DB, secretID string, secret string, ownerID string) error {
	err := execSQL(db, "INSERT INTO secrets (secretID, secret, ownerID) VALUES (?,?,?)", secretID, secret, ownerID)
	return err
}

func deleteSecret(db *sql.DB, secretID string) error {
	err := execSQL(db, "DELETE FROM secrets WHERE secretID = ?", secretID)
	return err
}

func updateSecret(db *sql.DB, secretID, newSecret string) error {
	err := execSQL(db, "UPDATE secrets SET secret = ? WHERE secretID = ?", newSecret, secretID)
	return err

}
func getSecret(db *sql.DB, secretID string) (Structures.Secret, string, error) {
	var secret Structures.Secret
	var ownerID string
	secret.SecretID = secretID
	err := db.QueryRow("SELECT secret, ownerID FROM secrets WHERE secretID = ?", secretID).Scan(&secret.Secret, &ownerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Structures.Secret{}, "", errNotFound
		}
		return Structures.Secret{}, "", errInternalServerError
	}
	return secret, ownerID, nil
}

func deleteUserByUsername(db *sql.DB, username string) error {
	err := execSQL(db, "DELETE FROM users WHERE username = ?", username)
	return err
}

func addUser(db *sql.DB, username string, password string) error {
	err := execSQL(db, "INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	return err
}

func hashSumGenerator(text string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), 14)
	if err != nil {
		return ""
	}
	return string(hash)
}
func getHeader(c *gin.Context, key string) string {
	h := c.GetHeader(key)
	return h
}
func getUserID(c *gin.Context) (string, bool) {
	userID, isExist := c.Get("user-id")
	return fmt.Sprint(userID), isExist
}
func updateUserPassword(db *sql.DB, username string, newPassword string) error {
	err := execSQL(db, "UPDATE users SET password = ? WHERE username = ?", newPassword, username)
	return err
}
func getUserIDByUsername(db *sql.DB, username string) (userID string, err error) {
	query := `SELECT userID FROM users WHERE username = ?`
	err = db.QueryRow(query, username).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errNotFound
		}
		return "", errInternalServerError
	}
	return userID, nil
}
