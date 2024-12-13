package Middleware

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/MustafaAbdulazizHamza/Pandora-Server/Structures"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	errUserNotFound         = errors.New("user not found")
	errAuthenticationFailed = errors.New("authentication failed")
	errInternalServerError  = errors.New("an internal server error occurred")
)

func abortWithError(c *gin.Context, statusCode int, message string) {
	response := Structures.Response{
		Status: fmt.Sprint(statusCode),
		Text:   message,
	}
	c.AbortWithStatusJSON(statusCode, response)
}
func getAuthData(c *gin.Context) (username string, password string, err error) {
	username = c.GetHeader("username")
	password = c.GetHeader("password")
	if password == "" || username == "" {
		return "", "", errAuthenticationFailed
	}
	return username, password, nil
}
func getUserInfoByUsername(db *sql.DB, username string) (userID string, password string, err error) {
	query := `SELECT password, userID FROM users WHERE username = ?`
	err = db.QueryRow(query, username).Scan(&password, &userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", errUserNotFound
		}
		return "", "", errInternalServerError
	}
	return userID, password, nil
}
func verifyPassword(hashedPassword string, plainPassword string) bool {
	// Convert the stored hash (string) to a byte slice for comparison
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	// Return true if the comparison succeeds, false otherwise
	return err == nil
}
func isCorrectLogin(password string, dbPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	return err == nil
}
func AuthenticateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		username, rpassword, err := getAuthData(c)
		if err != nil {
			abortWithError(c, http.StatusUnauthorized, err.Error())
			return
		}
		userID, password, err := getUserInfoByUsername(db, username)
		if err != nil {
			if errors.Is(err, errUserNotFound) {
				abortWithError(c, http.StatusUnauthorized, errAuthenticationFailed.Error())
				return
			}
			abortWithError(c, http.StatusInternalServerError, err.Error())
			return
		}

		if isAuthorized := isCorrectLogin(rpassword, password); !isAuthorized {
			abortWithError(c, http.StatusUnauthorized, errAuthenticationFailed.Error())
			return
		} else {
			c.Set("user-id", userID)
			c.Next()
		}

	}
}
