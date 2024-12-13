package APIs

import (
	"database/sql"
	"fmt"
	"github.com/MustafaAbdulazizHamza/Pandora/Structures"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostSecret(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			secretStructure Structures.Secret
			err             error
		)
		userID, isExist := getUserID(c)
		if !isExist {
			return
		}
		if err = c.ShouldBindJSON(&secretStructure); err != nil {
			response := Structures.Response{
				Status: fmt.Sprint(http.StatusBadRequest),
				Text:   "Unable to parse the secret you sent, please follow the typical structure of a secret.",
			}
			c.JSON(http.StatusBadRequest, response)
			return
		}
		if err = addSecret(db, secretStructure.SecretID, secretStructure.Secret, userID); err != nil {
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
