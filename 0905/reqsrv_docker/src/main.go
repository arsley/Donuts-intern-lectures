package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

func main() {
	e := echo.New()
	e.POST("/post", func(c echo.Context) error {
		return c.String(http.StatusOK, "post ok")
	})
	e.GET("/timeout", func(c echo.Context) error {
		time.Sleep(time.Minute)
		return c.String(http.StatusOK, "timeout ok")
	})
	e.GET("/status/:code", func(c echo.Context) error {
		codeStr := c.Param("code")
		code, err := strconv.Atoi(codeStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		codeText := http.StatusText(code)
		if codeText == "" {
			return c.String(http.StatusNotFound, "no such status code")
		}
		return c.String(code, codeText)
	})
	e.GET("/content/:type", func(c echo.Context) error {
		contentType := c.Param("type")
		switch contentType {
		case "json":
			return c.JSON(http.StatusOK, map[string]interface{}{
				"status_code": 200,
				"status":      "ok",
			})
		case "html":
			return c.HTML(http.StatusOK, `<html><title>title</title><body>ok</body></html>`)
		default:
			return c.String(http.StatusOK, "ok")
		}
	})
	e.Logger.Fatal(e.Start(":8080"))
}
