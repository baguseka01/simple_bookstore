package routes

import (
	"log"
	"os"

	"github.com/baguseka01/simple_bookstore/controllers"
	m "github.com/baguseka01/simple_bookstore/middleware"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() *echo.Echo {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}

	m.LogMiddleware(e)

	user := e.Group("/user")
	user.POST("/register", controllers.UserRegister)
	user.POST("/login", controllers.UserLogin)

	book := e.Group("/book")
	book.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))
	book.POST("/add", controllers.AddBook)
	book.PUT("/update/:id", controllers.UpdateBook)
	book.DELETE("/delete/:id", controllers.DeleteBook)
	book.GET("/detail/:id", controllers.DetailBook)
	book.GET("/all", controllers.AllBooks)

	return e
}
