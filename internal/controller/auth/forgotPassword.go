package auth

import (
	"fmt"
	"net/http"
	"net/smtp"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
)

// @Summary	Forgot Password
// @Tags		auth
// @Param		email	formData	string	true	"Email address"
// @Success	200		{object}	models.SuccessResponse
// @Failure	400		{object}	models.ErrorResponse
// @Router		/forgot-password [post]
func (pc *AuthController) ForgotPassword(c *gin.Context) {
	email := c.PostForm("email")

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

	password := GenerateRandomPassword(12)

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
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

	err = pc.UserRepository.SetUserPassword(c, string(encryptedPassword), int(user.ID))
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

	subject := "Сброс пароля"
	HTMLbody := fmt.Sprintf("<html><body><div style=\"margin-top: 50px;\"><p>Сбрасываем ваш пароль</p><p>Для авторизации используйте ваш email и пароль указанный ниже:</p><p><strong>%s</strong></p></div></body></html>", password)

	err = SendEmail(email, subject, HTMLbody, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_SEND_EMAIL",
					Message: "Couldn't send email to user",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: "Successfully send to email"})
}

func SendEmail(email string, subject string, HTMLbody string, c *gin.Context) error {
	to := []string{email}

	fromEmail := "nurzat.tynyshbekov.04@gmail.com"
	SMTPpassword := "wzxw klrm fpfo uxmz"
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	auth := smtp.PlainAuth("", fromEmail, SMTPpassword, host)

	msg := []byte(
		"From: <" + fromEmail + ">\r\n" +
			"To: " + email + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
			"\r\n" +
			HTMLbody)

	err := smtp.SendMail(address, auth, fromEmail, to, msg)

	if err != nil {
		return err
	}
	return nil
}
