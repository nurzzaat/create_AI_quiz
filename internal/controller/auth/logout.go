package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

//	@Summary	LogOut
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//
// @Security	ApiKeyAuth
//
//	@Success	200		{object}	models.SuccessResponse
//	@Failure	default	{object}	models.ErrorResponse
//	@Router		/logout [post]
func (lc *AuthController) Logout(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	splitToken := strings.Split(tokenString, " ")
	if len(splitToken) != 2 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_TOKEN",
					Message: "token doesn't provided",
				},
			},
		})
		return
	}

	token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(lc.Env.AccessTokenSecret), nil
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "INVALID_TOKEN",
					Message: "token isn't correct",
				},
			},
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "INVALID_TOKEN_CLAIMS",
					Message: "token isn't correct",
				},
			},
		})
		return
	}

	claims["exp"] = time.Now().Unix()

	c.JSON(http.StatusOK, models.SuccessResponse{Result: "Successfully logout"})
}
