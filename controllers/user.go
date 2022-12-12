package controllers

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/baguseka01/simple_bookstore/config"
	"github.com/baguseka01/simple_bookstore/middleware"
	"github.com/baguseka01/simple_bookstore/models"
	"github.com/labstack/echo/v4"
)

func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}

func UserRegister(e echo.Context) error {
	var data map[string]interface{}

	if err := e.Bind(&data); err != nil {
		return e.JSON(http.StatusBadRequest, "Error record when register")
	}

	if len(data["password"].(string)) < 6 {
		return e.JSON(http.StatusBadRequest, "Password more than 6 character")
	}

	if !ValidateEmail(strings.TrimSpace(data["email"].(string))) {
		return e.JSON(http.StatusInternalServerError, "Email wrong character example@gmail.com")
	}

	var user models.User
	config.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&user)
	if user.ID != 0 {
		return e.JSON(http.StatusInternalServerError, "Email is used")
	}

	newUser := models.User{
		FirtsName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}

	newUser.HashPassword(data["password"].(string))
	err := config.DB.Create(&newUser)
	if err != nil {
		log.Println("Fail to connect database")
	}

	return e.JSON(http.StatusCreated, &echo.Map{
		"message": "Succes register",
		"user":    newUser,
	})
}

func UserLogin(e echo.Context) error {
	var data map[string]string
	var user models.User

	if err := e.Bind(&data); err != nil {
		return e.JSON(http.StatusBadRequest, "Error record when login")
	}

	config.DB.Where("email=?", data["email"]).First(&user)
	if user.ID == 0 {
		return e.JSON(http.StatusInternalServerError, "Email not register")
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		return e.JSON(http.StatusInternalServerError, "Password is wrong")
	}

	token, err := middleware.GenerateJwtToken(user.ID, user.FirtsName)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Token is wrong")
	}

	return e.JSON(http.StatusOK, &echo.Map{
		"message": "Success login",
		"token":   token,
	})

}
