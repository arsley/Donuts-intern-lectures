package main

import (
	"fmt"
	"log"

	"github.com/flosch/pongo2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

var (
	templateSet *pongo2.TemplateSet = pongo2.DefaultSet
)

func init() {
	templateSet = pongo2.NewSet("default")
	templateSet.SetBaseDirectory("template")
}

// DB nice
type DB struct {
	*sqlx.DB
}

// store database instance
var (
	db *DB
)

func main() {
	// connect to db
	sqlxdb, err := sqlx.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		"minich_local_user",
		"minich_local_password",
		"db",
		"3306",
		"minich_local",
	))
	if err != nil {
		log.Fatalf("DB Connection Error: %v", err)
		return
	}
	db = &DB{sqlxdb}

	e := echo.New()
	e.Static("static", "static")

	e.GET("/movies", handleMovies)
	e.GET("/movie/:id", handleShowMovie)
	e.GET("/movie/:id/delete", handleDelete)
	e.GET("/movie/:id/edit", handleEdit)
	// e.POST("/movie/:id/update", handleUpdate)
	e.GET("/movies/new", handleNew)
	e.POST("/movies", handleUpload)
	e.Logger.Fatal(e.Start(":8080"))
}
