package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sahrialardians/rekeningku/internal/utils"
)

// Authenticated middleware untuk memverifikasi token JWT
func Authenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		// Validasi header
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.Unauthorized(ctx, "Authorization header is missing or invalid", nil)
			ctx.Abort()
			return
		}

		// Mengambil token dari header
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
		if tokenString == "" {
			utils.Unauthorized(ctx, "Token is missing or empty", nil)
			ctx.Abort()
			return
		}

		// Decode token
		userData, err := utils.DecryptJWT(tokenString)
		if err != nil {
			utils.Unauthorized(ctx, "Invalid token", nil)
			ctx.Abort()
			return
		}

		// Ambil user_id dari klaim
		userID, _ := userData["user_id"].(float64)

		// Simpan user_id ke context dengan key "user_id"
		ctx.Set("user_id", int(userID))
		ctx.Next()
	}
}
