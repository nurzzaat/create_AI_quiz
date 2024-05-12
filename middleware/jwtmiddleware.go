package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/controller/tokenutil"
	models "github.com/nurzzaat/create_AI_quiz/internal/models"
)

func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(1)
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
		fmt.Println(2)
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
