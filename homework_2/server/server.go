package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"homework_2/handler"
)

func main() {
	accountHandler := handler.New()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Queries
	e.POST("account/create", accountHandler.CreateAccount)
	e.DELETE("account/delete", accountHandler.DeleteAccount)
	e.POST("account/update/amount", accountHandler.UpdateAmount)
	e.PUT("account/update/name", accountHandler.UpdateName)
	e.GET("account", accountHandler.GetAccount)

	// Launch
	e.Logger.Fatal(e.Start(":1323"))
}
