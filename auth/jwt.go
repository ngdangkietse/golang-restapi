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

type JwtUserClaims struct {
	UserId int    `json:"userId"`
	RoleId int    `json:"roleId"`
	Email  string `json:"email"`
}

func GenerateToken(user models.User) (string, error) {
	tokenExp, _ := strconv.Atoi(os.Getenv("TOKEN_EXP"))

	exp := time.Now().Add(time.Duration(tokenExp) * time.Hour).Unix()

	mapClaims := make(jwt.MapClaims)
	mapClaims["sub"] = user.Email
	mapClaims["user"] = JwtUserClaims{
		UserId: user.Id,
		RoleId: user.RoleId,
		Email:  user.Email,
	}
	mapClaims["iat"] = time.Now().Unix()
	mapClaims["nbf"] = time.Now().Unix()
	mapClaims["exp"] = exp

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims).SignedString(jwtSecretKey)

	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	return token, nil
}

func ValidateToken(jwtToken string) (interface{}, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !token.Valid || !ok {
		return nil, fmt.Errorf(err.Error())
	}

	return claims["user"], nil
}

func GetTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
