package common

import (
	"fmt"
	"go-transaction/config"
	"go-transaction/model"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func IDFromToken(ctx *gin.Context) string {
	var header model.AuthHeader
	ctx.ShouldBindHeader(&header)

	tokenString := strings.Replace(header.AuthorizationHeader, "Bearer ", "", 1)

	cfg := config.NewConfig()
	token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return cfg.JwtSignatureKey, nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)
	IdToken := fmt.Sprintf("%v", claims["Id"])
	return IdToken
}
