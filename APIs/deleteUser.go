package APIs

import (
	"database/sql"
	"fmt"
	"github.com/MustafaAbdulazizHamza/Pandora-Server/Structures"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := getUserID(c)
		if !exists {
			return
		}

		if userID == "0" {
			var (
				usr Structures.User
				err error
			)
			if err = c.ShouldBindJSON(&usr); err != nil {
				response := Structures.Response{
					Status: fmt.Sprint(http.StatusBadRequest),
					Text:   err.Error(),
				}
				c.JSON(http.StatusBadRequest, response)
				return
			}
			if err = deleteUserByUsername(db, usr.Username); err != nil {
				response := Structures.Response{
					Status: fmt.Sprint(http.StatusInternalServerError),
					Text:   err.Error(),
				}
				c.JSON(http.StatusInternalServerError, response)
				return
			}
			response := Structures.Response{
				Status: fmt.Sprint(http.StatusOK),
				Text:   "",
			}
			c.JSON(http.StatusOK, response)
		} else {
			response := Structures.Response{
				Status: fmt.Sprint(http.StatusUnauthorized),
				Text:   "You must be the root to carry out this process.",
			}

			c.JSON(http.StatusForbidden, response)

		}
	}
}
