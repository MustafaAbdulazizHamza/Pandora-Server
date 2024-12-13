package APIs

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/MustafaAbdulazizHamza/Pandora-Server/Structures"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateUserCredentials(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := getUserID(c)
		if !exists {
			return
		}
		var (
			userStructure Structures.User
			err           error
		)
		if err = c.ShouldBindJSON(&userStructure); err != nil {
			response := Structures.Response{
				Status: fmt.Sprint(http.StatusBadRequest),
				Text:   err.Error(),
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		ruserID, err := getUserIDByUsername(db, userStructure.Username)
		if err != nil {
			var s int
			if errors.Is(err, errInternalServerError) {
				s = http.StatusInternalServerError
			} else if errors.Is(err, errNotFound) {
				s = http.StatusNotFound
			}
			response := Structures.Response{
				Status: fmt.Sprint(s),
				Text:   err.Error(),
			}
			c.JSON(s, response)
		}
		if (ruserID != userID) && (userID != "0") {
			response := Structures.Response{Status: fmt.Sprint(http.StatusForbidden), Text: "You are not allowed to carry out this process."}
			c.JSON(http.StatusForbidden, response)
			return
		}
		userStructure.Password = hashSumGenerator(userStructure.Password)
		if err = updateUserPassword(db, userStructure.Username, userStructure.Password); err != nil {
			response := Structures.Response{
				Status: fmt.Sprint(http.StatusInternalServerError),
				Text:   errInternalServerError.Error(),
			}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
		response := Structures.Response{Status: fmt.Sprint(http.StatusOK), Text: ""}
		c.JSON(http.StatusOK, response)

	}
}
