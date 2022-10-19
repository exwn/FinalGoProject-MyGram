package middlewares

import (
	"MyGram/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if tokenData, err := helpers.VerifyToken(ctx); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		} else {
			ctx.Set("userData", tokenData)
			ctx.Next()
		}
	}
}
