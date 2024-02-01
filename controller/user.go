package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
func (h *UserHandler) GetUserById(c *gin.Context) {
	var user database.User
	id := c.Param("id")

	h.DB.Where("id = ?", id).First(&user)
	c.JSON(http.StatusOK, response.NewResponse(200, "Successfully get Users Data", user))
}

func (h *UserHandler) AddUser(c *gin.Context) {
	var users database.User
	c.Bind(&users)
	validate := validator.New()

	err := validate.Struct(users)
	if err != nil {
		http.Error(c.Writer, "All fields must not be empty", http.StatusBadRequest)
		return
	}

	if users.Password != users.ConfPassword {
		http.Error(c.Writer, "Password and Confirm Password do not match", http.StatusBadRequest)
		return
	}

	hash, _ := HashPassword(users.Password)
	users.Password = hash

	result := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		Name:  users.Name,
		Email: users.Email,
	}

	// Tambahkan user ke dalam database
	h.DB.Create(&users)

	// Mengembalikan response JSON
	c.JSON(http.StatusOK, gin.H{"message": "Successfully add user", "user": result})
}
