package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, ok := ctx.Get("claims")
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "invalid token",
			})
			ctx.Abort()
			return
		}

		claimsMap := claims.(jwt.MapClaims)

		if claimsMap["Role"] != "admin" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "access denied",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
