package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/baguseka01/simple_bookstore/config"
	"github.com/baguseka01/simple_bookstore/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var DB *gorm.DB

func AddBook(e echo.Context) error {
	var book models.Book
	if err := e.Bind(&book); err != nil {
		return e.JSON(http.StatusBadRequest, "Error record book")
	}

	err := config.DB.Create(&book).Error
	if err != nil {
		log.Println("Not connect database create book")
	}

	return e.JSON(http.StatusOK, &echo.Map{
		"message": "Succes add book",
	})
}

func UpdateBook(e echo.Context) error {
	var book models.Book
	if err := e.Bind(&book); err != nil {
		return e.JSON(http.StatusBadRequest, "Error record book")
	}

	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Where("id=?", id).Updates(&book).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, &echo.Map{
			"message": "Can't update book",
			"error":   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"message":     "Success update book",
		"update book": book,
	})
}

func DeleteBook(e echo.Context) error {
	var book models.Book

	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Model(&book).Where("id=?", id).Delete(&book).Error
	if err != nil {
		log.Println("Not connect database delete book")
	}

	return e.JSON(http.StatusOK, &echo.Map{
		"message": "Success delete book",
	})
}

func DetailBook(e echo.Context) error {
	var book models.Book

	id, _ := strconv.Atoi(e.Param("id"))
	err := config.DB.Where("id=?", id).Preload("Admin").First(&book).Error
	if err != nil {
		log.Println("Not connect database detail book")
	}

	return e.JSON(http.StatusOK, &echo.Map{
		"book": book,
	})
}

func AllBooks(e echo.Context) error {
	var books []models.Book

	err := config.DB.Model(&books).Find(&books).Error
	if err != nil {
		log.Println("Not connect database all book")
	}

	return e.JSON(http.StatusOK, &echo.Map{
		"books": books,
	})
}
