package main

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"github.com/nelsonstr/monorepo/libs/hello"
)

func main() {
	e := echo.New()
	e.GET("/two/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, hello.Greet("Friend2"))
	})
	_ = e.Start(":8080")
}
