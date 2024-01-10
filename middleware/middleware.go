package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ordinary-fdev/go-learning/services"
)

func ValidateRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := services.VerifyToken(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
