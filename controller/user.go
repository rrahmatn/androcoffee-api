package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rrahmatn/androcoffee-api.git/database"
	"github.com/rrahmatn/androcoffee-api.git/response"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

// GetUsers handles GET request to fetch all users
func (h *UserHandler) GetUsers(c *gin.Context) {
	var users []database.User
	h.DB.Find(&users)
	c.JSON(http.StatusOK, response.NewResponse(200, "Successfully get Users Data", users))
}

func (h *UserHandler) AddUser(c *gin.Context) {
	var users database.User
	c.Bind(&users)
	hash, _ := HashPassword(users.Password)

	users.Password = hash

	h.DB.Create(&users)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully add user", "user": users})
}
