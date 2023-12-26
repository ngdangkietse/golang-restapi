package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang-rest-api/models"
	"os"
	"strconv"
	"strings"
	"time"
)

var jwtSecretKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateToken(email string, userId int, roleId int) (*models.LoginResponse, error) {
	tokenExp, _ := strconv.Atoi(os.Getenv("TOKEN_EXP"))

	exp := time.Now().Add(time.Duration(tokenExp) * time.Hour).Unix()

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"roleId": roleId,
		"email":  email,
		"iat":    time.Now().Unix(),
		"exp":    exp,
	})

	token, err := claims.SignedString(jwtSecretKey)

	return &models.LoginResponse{
		Token:  token,
		UserId: userId,
		RoleId: roleId,
	}, err
}

func ValidateToken(jwtToken string) (bool, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		return false, err
	}

	if token.Valid {
		return true, nil
	}

	return false, err
}

func GetTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
