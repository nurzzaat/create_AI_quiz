package auth

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

// @Summary	Account Recovery
// @Tags		auth
// @Param		email			path		string	true	"User email"
// @Param		hash_pass		path		string	true	"User's hashed password"
// @Param		password		formData	string	true	"New password"
// @Param		confirmPassword	formData	string	true	"Confirm new password"
// @Success	200				{object}	models.SuccessResponse
// @Failure	400				{object}	models.ErrorResponse
// @Router		/accountrecovery/{email}/{hash_pass} [post]
func (pc *AuthController) AccountRecovery(c *gin.Context) {
	email := c.Param("email")
	password := c.Param("hash_pass")

	user, err := pc.UserRepository.GetUserByEmail(c, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_USER",
					Message: "User with this email doesn't found",
				},
			},
		})
		return
	}

	if user.Password != password {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "PASSWORDS_DONT_MATCH",
					Message: "Passwords from database aren't the same",
				},
			},
		})
		return
	}
	password1 := c.PostForm("password")
	password2 := c.PostForm("confirmPassword")
	if password1 != password2 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "PASSWORDS_DONT_MATCH",
					Message: "The passwords aren't the same",
					Metadata: models.Properties{
						Properties1: password1,
						Properties2: password2,
					},
				},
			},
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.MinCost)
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

	user.Password = string(hash)
	err = pc.UserRepository.SetUserPassword(c, user.Password, int(user.ID))
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

	c.HTML(http.StatusOK, "reset-password.html", gin.H{"message": "Ваш пароль обнавлен!"})
}

func (pc *AuthController) AccountRecoveryHTML(c *gin.Context) {
	email := c.Param("email")
	hash := c.Param("hash_pass")
	c.HTML(http.StatusOK, "reset-password.html", gin.H{"Email": email, "Hash": hash})
}
