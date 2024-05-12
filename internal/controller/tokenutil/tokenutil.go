package tokenutil

import (
	"fmt"
	"time"

	models "github.com/nurzzaat/create_AI_quiz/internal/models"

	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error) {
	//exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &models.JwtClaims{
		ID:     user.ID,
		RoleID: user.RoleID,
		// RegisteredClaims: jwt.RegisteredClaims{
		// 	ExpiresAt: jwt.NewNumericDate(exp),
		// },
	}
	fmt.Println("claim", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error) {
	claimsRefresh := &models.JwtRefreshClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return rt, err
}

func ValidateJWT(c *gin.Context, secret string) error {
	fmt.Println(11)
	token, err := getToken(c, secret)
	if err != nil {
		return err
	}
	fmt.Println(12)
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func ValidateUserJWT(c *gin.Context, secret string) error {
	fmt.Println(21)
	token, err := getToken(c, secret)
	if err != nil {
		return err
	}
	fmt.Println(22)
	claims, ok := token.Claims.(jwt.MapClaims)
	userRoleID := uint(claims["role"].(float64))
	userID := uint(claims["id"].(float64))
	if ok && token.Valid {
		c.Set("userID", userID)
		c.Set("roleID", userRoleID)
		return nil
	}
	return errors.New("invalid curator token provided")
}

func getToken(c *gin.Context, secret string) (*jwt.Token, error) {
	fmt.Println(111)
	tokenString := getTokenFromRequest(c)
	fmt.Println(112)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		fmt.Println(113)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	return token, err
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	fmt.Println(114)
	fmt.Println(splitToken)
	fmt.Println(bearerToken)
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	fmt.Println(1115)
	return ""
}
