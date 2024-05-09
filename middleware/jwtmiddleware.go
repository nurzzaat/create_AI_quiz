package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/ZharasDiplom/internal/controller/tokenutil"
	models "github.com/nurzzaat/ZharasDiplom/internal/models"
)

func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := tokenutil.ValidateJWT(c, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Result: []models.ErrorDetail{
					{
						Code:    "Authorization error",
						Message: "Authorization error",
						Metadata: models.Properties{
							Properties1: err.Error(),
						},
					},
				},
			})
			c.Abort()
			return
		}
		err = tokenutil.ValidateUserJWT(c, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Result: []models.ErrorDetail{
					{
						Code:    "User is required",
						Message: "User is required",
						Metadata: models.Properties{
							Properties1: err.Error(),
						},
					},
				},
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
