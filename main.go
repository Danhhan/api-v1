package main

import (
	// "fmt"
	"go-module/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", welcome)
	e.GET("/user/sign-in", routes.HandleSignIn)
	e.GET("/user/sign-up", routes.HandleSignUp)

	e.Logger.Fatal(e.Start(":3000"))
}
func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcom to my api!")
}