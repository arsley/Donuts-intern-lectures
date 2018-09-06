package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/v2/:path", func(c echo.Context) error {
		return c.String(http.StatusOK, echo.Context.Param(c, "path"))
	})
	e.Logger.Fatal(e.Start(":80"))
}
