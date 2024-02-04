package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rrahmatn/androcoffee-api.git/database"
	"github.com/rrahmatn/androcoffee-api.git/requests"
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

func (h *UserHandler) GetUsers(c *gin.Context) {
	var users []database.User
	h.DB.Find(&users)

	c.JSON(http.StatusOK, response.NewResponse(200, "Successfully get Users Data", users))
}
func (h *UserHandler) GetUserById(c *gin.Context) {
	var users database.User
	id := c.Param("id")

	h.DB.Where("id = ?", id).First(&users)
	c.JSON(http.StatusOK, response.NewResponse(200, "Successfully get Users Data", users))
}

func (h *UserHandler) AddUser(c *gin.Context) {
	var newUser requests.CreateUser
	c.Bind(&newUser)
	validate := validator.New()

	err := validate.Struct(newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "All field must not be empty"})
		return
	}

	if newUser.Password != newUser.ConfPassword {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Password and Confirm Password must match"})
		return
	}

	var users database.User
	result := h.DB.Where("email = ?", newUser.Email).First(&users)
	if result.RowsAffected > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Email already exists"})
		return
	}

	hash, _ := HashPassword(newUser.Password)

	user := database.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: hash,
	}

	h.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully add user", "user": user})
}
