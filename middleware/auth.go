package middleware

import (
	"database/sql"
	"net/http"
	"os"
	"quiz-3/helpers"
	"quiz-3/repository"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var (
	DB *sql.DB
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			helpers.ResponseJSON(c, http.StatusUnauthorized, "Missing Token", "error", nil)

			return
		}

		if !strings.HasPrefix(tokenString, "Bearer ") {
			helpers.ResponseJSON(c, http.StatusUnauthorized, "Invalid Token Format", "error", nil)
			return
		}

		tokenString = tokenString[7:]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			helpers.ResponseJSON(c, http.StatusUnauthorized, "Invalid Token", "error", nil)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			helpers.ResponseJSON(c, http.StatusInternalServerError, "Internal Server Error", "error", nil)

			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			helpers.ResponseJSON(c, http.StatusInternalServerError, "Internal Server Error", "error", nil)

			return
		}

		userID := int(userIDFloat)

		userRepo := &repository.UserRepository{DB: DB}

		username, err := userRepo.GetUserByID(userID)
		if err != nil {
			helpers.ResponseJSON(c, http.StatusInternalServerError, "Internal Server Error", "error", nil)

			return
		}

		c.Set("user", username.Username)

		c.Next()
	}
}
