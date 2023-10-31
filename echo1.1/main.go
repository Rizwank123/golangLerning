package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello and Welcome")

	})
	e.Logger.Print("listing on port 8000")
	e.Logger.Fatal(e.Start(":8000"))

}
