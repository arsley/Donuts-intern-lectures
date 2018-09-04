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
	e.GET("/", handlePage)
	e.GET("/:page", handlePage)
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8080"))
}

func handlePage(c echo.Context) error {
	page := echo.Context.Param(c, "page")
	if page == "" {
		page = "index"
	}
	page = "template/" + page + ".html"

	html, err := ioutil.ReadFile(page)
	if _, ok := err.(*os.PathError); ok {
		return c.HTML(http.StatusNotFound, "<h1>File not found</h1>")
	} else if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusInternalServerError, "<h1>Internal server error</h1>")
	}

	return c.HTMLBlob(http.StatusOK, html)
}
