package handlers

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/wpcodevo/two_factor_golang/models"

	// "github.com/pquerna/otp/totp"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

func (ac *AuthController) SignUpUser(c echo.Context) error {
	var payload *models.RegisterUserInput

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "fail", "message": err.Error()})
	}

	newUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: payload.Password,
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.JSON(http.StatusConflict, map[string]interface{}{"status": "fail", "message": "Email already exists, please use another email address"})
	} else if result.Error != nil {
		return c.JSON(http.StatusBadGateway, map[string]interface{}{"status": "error", "message": result.Error.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"status": "success", "message": "Registered successfully, please login"})
}

func (ac *AuthController) LoginUser(c echo.Context) error {
	var payload *models.LoginUserInput

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "fail", "message": err.Error()})
	}

	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": "fail", "message": "Invalid email or Password"})
	}

	userResponse := map[string]interface{}{
		"id":          user.ID.String(),
		"name":        user.Name,
		"email":       user.Email,
		"otp_enabled": user.Otp_enabled,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"status": "success", "user": userResponse})
}
