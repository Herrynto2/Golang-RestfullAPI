package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware ...
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateToken(c)
		c.Next()
	}
}

func validateToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if len(token) == 0 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "You cannot access this API!",
		})
	} else {
		token = strings.Split(token, "Bearer ")[1]
		userID, err := getPayload(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "You cannot access this API!",
			})
		}
		c.Set("userID", userID)
		c.Next()
	}
}

func getPayload(tokenString string) (interface{}, error) {
	atClaims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, atClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		// panic(err)
		return nil, err
	}
	return atClaims["user_id"], nil
}
