package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nelsonstr/monorepo/libs/hello"
)

func main() {
	e := echo.New()
	e.GET("/one/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, hello.Greet("World"))
	})
	_ = e.Start(":8080")
}
