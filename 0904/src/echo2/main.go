package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", handlerHello)
	e.GET("/:name", handlerHello)
	e.Logger.Fatal(e.Start(":8080"))
}

func handlerHello(c echo.Context) error {
	html := `
        <!DOCTYPE html>
        <html>
            <h1>Hello, world</h1>
        </html>
    `
	return c.HTML(http.StatusOK, html)
}
