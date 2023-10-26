package middleware

import (
	"go-transaction/model"
	"go-transaction/utils/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var header model.AuthHeader
		err := ctx.ShouldBindHeader(&header)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "invalid header " + err.Error(),
			})
			ctx.Abort()
			return
		}

		tokenString := strings.Replace(header.AuthorizationHeader, "Bearer ", "", 1)

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "unauthorized",
			})
			ctx.Abort()
			return
		}

		claims, err := security.VerifyAccessToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"message": "invalid token " + err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
