package api

import (
	"github.com/YounesBouchbouk/natours-api-golang/token"
	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func AuthenticationMiddlware(server Server) gin.HandlerFunc {

	maker, err := token.NewJWTMaker(server.config.JWT_secret_key)

	if err != nil {
		return func(ctx *gin.Context) {
			respondWithError(ctx, 400, "token is required")
		}
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Token")
		if token == "" {
			respondWithError(ctx, 401, "API token required")
			return
		}
		payload, err := maker.VerifyToken(token)
		if err != nil {
			respondWithError(ctx, 400, "token is invalid")
			return
		}
		ctx.Set("email", payload.Email)
		ctx.Set("role", payload.Role)
		ctx.Next()
	}
}

func CheckIfAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, _ := ctx.Get("role")
		if role != "admin" {
			respondWithError(ctx, 401, "permission denied")
			return
		}
		ctx.Next()
	}
}
