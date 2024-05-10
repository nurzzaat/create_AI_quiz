package auth

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

// @Summary	Reset password
// @Security	ApiKeyAuth
// @Tags		auth
// @Accept		json
// @Produce	json
// @Param		reset	body		models.Password	true	"Change password"
// @Success	200		{object}	models.SuccessResponse
// @Failure	400		{object}	models.ErrorResponse
// @Router		/user/reset-password [post]
func (pc *AuthController) ResetPassword(c *gin.Context) {
	userID := c.GetUint("userID")
	var passwords models.Password

	err := c.BindJSON(&passwords)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of password",
				},
			},
		})
		return
	}

	if passwords.ConfirmPassword == "" || passwords.Password == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "EMPTY_FIELDS",
					Message: "Required fields are missing or invalid",
				},
			},
		})
		return
	}

	if passwords.Password != passwords.ConfirmPassword {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "PASSWORDS_DONT_MATCH",
					Message: "The passwords aren't the same",
					Metadata: models.Properties{
						Properties1: passwords.Password,
						Properties2: passwords.ConfirmPassword,
					},
				},
			},
		})
		return
	}

	err = validatePassword(passwords.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_PASSWORD_FORMAT",
					Message: err.Error(),
				},
			},
		})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(passwords.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_ENCRYPTE_PASSWORD",
					Message: "Couldn't encrypte password",
				},
			},
		})
		return
	}

	passwords.Password = string(encryptedPassword)

	err = pc.UserRepository.SetUserPassword(c, passwords.Password, int(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_RESET_PASSWORD",
					Message: "Couldn't reset password",
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: "Successfully updated password"})
}
