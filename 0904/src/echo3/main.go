package main

import (
	"net/http"

	"github.com/MakeNowJust/heredoc"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", handlerHello)
	e.GET("/:name", handlerHello)
	e.Logger.Fatal(e.Start(":8080"))
}

func handlerHello(c echo.Context) error {
	name := echo.Context.Param(c, "name")
	if name == "" {
		name = "Bob"
	}

	html := heredoc.Docf(`
        <!DOCTYPE html>
        <html>
            <h1>Hello, %s</h1>
        </html>
    `, name)
	return c.HTML(http.StatusOK, html)
}
