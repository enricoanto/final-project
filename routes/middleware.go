package routes

import (
	"errors"
	"net/http"
	"strings"

	"github.com/enricoanto/final-project/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func ParseToken(token string) (*jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("Rahasia"), nil
	})

	return jwtToken, claims, err
}

func GetTokenHeader(c *gin.Context) string {
	sign := strings.Split(c.Request.Header.Get("Authorization"), " ")
	return sign[1]
}

func(controller *Controller) Middleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := Claims(c)
		if err != nil {
			helper.Error(c, http.StatusUnauthorized, errors.New("invalid token"))
			c.Abort()
			return
		}

		if user["role"].(string) != "admin" || role != "admin" {
			helper.Error(c, http.StatusUnauthorized, errors.New("access denied"))
			c.Abort()
			return
		}

		if role == "customer" {
			c.Next()
		}
		c.Next()
	}
}

func Claims(c *gin.Context) (map[string]interface{}, error) {
	token := GetTokenHeader(c)

	_, claims, err := ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
