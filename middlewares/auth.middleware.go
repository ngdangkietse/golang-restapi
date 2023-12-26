package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-rest-api/auth"
	"golang-rest-api/payload"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := auth.GetTokenFromRequest(c)
		if token == "" {
			c.IndentedJSON(http.StatusUnauthorized, payload.HandleException("Request not contains access_token"))
			return
		}

		_, err := auth.ValidateToken(token)

		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, payload.HandleException(fmt.Sprintf("Error: [%v]", err.Error())))
			return
		}

		c.Next()
	}
}
