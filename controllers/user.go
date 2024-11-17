package controllers

import (
	"net/http"
	"os"
	"quiz-3/helpers"
	"quiz-3/repository"
	"quiz-3/structs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Login godoc
// @Summary User login
// @Description Logs in a user and generates a JWT token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param login body structs.UserPayload true "User login credentials"
// @Success 200 {object} helpers.Response "Login successful"
// @Failure 400 {object} helpers.Response "Invalid input data"
// @Failure 401 {object} helpers.Response "Invalid username or password"
// @Failure 500 {object} helpers.Response "Failed to generate token"
// @Router /api/login [post]
func Login(repo *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload structs.UserPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			helpers.ResponseJSON(c, http.StatusBadRequest, "invalid input data", "error", nil)

			return
		}

		var userID int
		var hashedPassword string
		query := "SELECT id, password FROM users WHERE username = $1"
		err := repo.DB.QueryRow(query, payload.Username).Scan(&userID, &hashedPassword)
		if err != nil {
			helpers.ResponseJSON(c, http.StatusUnauthorized, "invalid username or password", "error", nil)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(payload.Password)); err != nil {
			helpers.ResponseJSON(c, http.StatusUnauthorized, "invalid username or password", "error", nil)
			return
		}

		token, err := GenerateJWT(userID)
		if err != nil {
			helpers.ResponseJSON(c, http.StatusInternalServerError, "failed to generate token", "error", nil)
			return
		}

		helpers.ResponseJSON(c, http.StatusOK, "Login Success", "success", token)
	}
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Registers a new user and stores the user data in the database
// @Tags auth
// @Accept  json
// @Produce  json
// @Param register body structs.UserPayload true "User registration data"
// @Success 200 {object} helpers.Response "User registered successfully"
// @Failure 400 {object} helpers.Response "Invalid input data"
// @Failure 409 {object} helpers.Response "Username already exists"
// @Failure 500 {object} helpers.Response "Failed to register user"
// @Router /api/register [post]
func RegisterUser(repo *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req structs.UserPayload

		// Bind JSON ke struct
		if err := c.ShouldBindJSON(&req); err != nil {
			helpers.ResponseJSON(c, http.StatusBadRequest, "invalid input data", "error", nil)
			return
		}

		// Hash password menggunakan bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			helpers.ResponseJSON(c, http.StatusInternalServerError, "failed to hash password", "error", nil)
			return
		}

		// Simpan user ke database melalui metode InsertUser
		err = repo.InsertUser(req.Username, string(hashedPassword))
		if err != nil {
			if err.Error() == "username already exists" {
				helpers.ResponseJSON(c, http.StatusConflict, "username already exists", "error", nil)
				return
			}
			helpers.ResponseJSON(c, http.StatusInternalServerError, "failed to register user", "error", nil)
			return
		}

		// Response berhasil
		helpers.ResponseJSON(c, http.StatusOK, "user registered successfully", "message", nil)
	}
}
