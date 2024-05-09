package auth

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"unicode"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/gin-gonic/gin"
	"github.com/nurzzaat/create_AI_quiz/internal/controller/tokenutil"
	"github.com/nurzzaat/create_AI_quiz/internal/models"
	"github.com/nurzzaat/create_AI_quiz/pkg"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	UserRepository models.UserRepository
	Env            *pkg.Env
}

var (
	verifier = emailverifier.NewVerifier()
)

// @Summary	SignUp
// @Tags		auth
// @Accept		json
// @Produce	json
// @Param		user	body		models.UserRequest	true	"user"
// @Success	200		{object}	models.SuccessResponse
// @Failure	default	{object}	models.ErrorResponse
// @Router		/signup [post]
func (uc *AuthController) Signup(c *gin.Context) {
	var request models.UserRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of signup",
				},
			},
		})
		return
	}

	verifier = verifier.EnableSMTPCheck()
	verifier = verifier.EnableDomainSuggest()

	if request.Email == "" || request.Password == "" || request.FirstName == "" || request.LastName == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "EMPTY_FIELDS",
					Message: "Not all fields provided",
				},
			},
		})
		return
	}

	ret, _ := verifier.Verify(request.Email)

	log.Println(ret.Reachable)

	if !ret.Syntax.Valid {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "Неверный адрес электронной почты",
					Message: "email address syntax is invalid",
				},
			},
		})
		return
	}
	if ret.Reachable != "yes" && !isICloudEmail(request.Email) {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "Неверный адрес электронной почты",
					Message: "email address is not reachable",
				},
			},
		})
		return
	}

	user, _ := uc.UserRepository.GetUserByEmail(c, request.Email)
	if user.ID > 0 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "USER_EXISTS",
					Message: "User with this email already exists",
				},
			},
		})
		return
	}
	err := validatePassword(request.Password)
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
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
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
	request.Password = string(encryptedPassword)

	_, err = uc.UserRepository.CreateUser(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_CREATE_USERS",
					Message: "Couldn't create user",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}
	user, err = uc.UserRepository.GetUserByEmail(c, request.Email)
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
	accessToken, err := tokenutil.CreateAccessToken(&user, uc.Env.AccessTokenSecret, uc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "TOKEN_ERROR",
					Message: "Error to create access token",
				},
			},
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessResponse{Result: accessToken})
}

func isICloudEmail(email string) bool {
	icloudPattern := `@icloud\.com$`
	icloudRegex := regexp.MustCompile(icloudPattern)

	return icloudRegex.MatchString(email)
}

func GenerateRandomPassword(size int) string {
	var alpha = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	password := make([]rune, size)
	for i := 0; i < size; i++ {
		password[i] = alpha[rand.Intn(len(alpha)-1)]
	}
	hashPassword := string(password)
	return hashPassword
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	var (
		hasUpper, hasLower, hasDigit bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
			// case unicode.IsPunct(char) || unicode.IsSymbol(char):
			// 	hasSpecial = true
		}
	}
	if !hasUpper || !hasLower || !hasDigit {
		return fmt.Errorf("password must contain at least one uppercase letter, one lowercase letter and one digit")
	}
	return nil
}
