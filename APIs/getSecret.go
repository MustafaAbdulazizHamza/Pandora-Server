package APIs

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/MustafaAbdulazizHamza/Pandora/Structures"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSecret(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			secretStructure Structures.RequestedSecret
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

		dbSECRET, ownerID, err := getSecret(db, secretStructure.SecretID)
		if err != nil {
			s := http.StatusInternalServerError
			if errors.Is(err, errNotFound) {
				s = http.StatusNotFound
			}
			response := Structures.Response{
				Status: fmt.Sprint(s),
				Text:   err.Error(),
			}
			c.JSON(s, response)
			return
		}
		if ownerID != userID {
			response := Structures.Response{
				Status: fmt.Sprint(http.StatusForbidden),
				Text:   "You are not allowed to access the requested secret.",
			}
			c.JSON(http.StatusForbidden, response)
			return
		}
		response := Structures.Response{Status: fmt.Sprint(http.StatusOK), Text: dbSECRET.Secret}
		c.JSON(http.StatusOK, response)
	}
}
