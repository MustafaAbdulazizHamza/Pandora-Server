package APIs

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/MustafaAbdulazizHamza/Pandora-Server/Structures"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateSecret(db *sql.DB) gin.HandlerFunc {
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
		_, ownerID, err := getSecret(db, secretStructure.SecretID)
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
				Text:   "You are not permitted to alter this secret.",
			}
			c.JSON(http.StatusForbidden, response)
			return

		}
		if err = updateSecret(db, secretStructure.SecretID, secretStructure.Secret); err != nil {
			fmt.Println(err)
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
