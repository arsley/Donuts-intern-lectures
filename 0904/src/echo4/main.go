package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", handlerHello)
	e.Logger.Fatal(e.Start(":8080"))
}

func handlerHello(c echo.Context) error {
	html, err := ioutil.ReadFile("../template/hello.html")
	if _, ok := err.(*os.PathError); ok {
		return c.HTML(http.StatusNotFound, "<h1>File not found</h1>")
	} else if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusInternalServerError, "<h1>Internal server error</h1>")
	}

	return c.HTMLBlob(http.StatusOK, html)
}
