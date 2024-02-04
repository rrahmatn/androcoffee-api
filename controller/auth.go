package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/rrahmatn/androcoffee-api.git/database"
	"github.com/rrahmatn/androcoffee-api.git/requests"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

func CreateToken(id int, name string, email string) (string, error) {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		return "Error Loading .env file", err
	}

	secret := []byte(myEnv["ACCESS_TOKEN"])

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    id,
			"name":  name,
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "Ini error saat membuat token", err
	}

	return tokenString, nil
}
func RefreshToken(id int, name string, email string) (string, error) {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		return "Error Loading .env file", err
	}

	secret := []byte(myEnv["REFRESH_TOKEN"])

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    id,
			"name":  name,
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "Ini error saat membuat token", err
	}

	return tokenString, nil
}

func (h *AuthHandler) Signin(c *gin.Context) {
	var loginData requests.Login

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if loginData.Email == "" || loginData.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	var user database.User
	result := h.DB.Where("email = ?", loginData.Email).First(&user)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Email not registered"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Wrong Password"})
		return
	}

	at, err := CreateToken(int(user.Id), user.Name, user.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	rt, err := RefreshToken(int(user.Id), user.Name, user.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": at, "refresh_token": rt})
}
