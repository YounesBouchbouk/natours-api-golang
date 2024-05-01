package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/YounesBouchbouk/natours-api-golang/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "autorization-payload-key"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func AuthenticationMiddlware(server Server) gin.HandlerFunc {

	// to is used to check if provided token is valid , line 29
	maker, err := token.NewJWTMaker(server.config.JWT_secret_key)

	if err != nil {
		return func(ctx *gin.Context) {
			respondWithError(ctx, 400, "Some thing went wrong")
		}
	}

	return func(ctx *gin.Context) {

		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			respondWithError(ctx, http.StatusUnauthorized, "API token required")
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) != 2 {
			respondWithError(ctx, http.StatusUnauthorized, "authorization Header format is invalid")

		}
		authorizationType := strings.ToLower(fields[0])

		if authorizationTypeBearer != authorizationType {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			respondWithError(ctx, http.StatusUnauthorized, err)
		}

		token := fields[1]
		payload, err := maker.VerifyToken(token)
		if err != nil {
			respondWithError(ctx, http.StatusUnauthorized, "token is invalid")
			return
		}
		ctx.Set(authorizationPayloadKey, payload)
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
