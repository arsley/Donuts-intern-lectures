package main

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
)

func handleMovies(c echo.Context) error {
	movies := db.FetchMovies()
	return render(c, "index.html", map[string]interface{}{"movies": movies})
}

func handleShowMovie(c echo.Context) error {
	id, _ := strconv.Atoi(echo.Context.Param(c, "id"))
	movie := db.FetchMovieByID(int64(id))
	return render(c, "show.html", pongo2.Context{"movie": movie})
}

func handleNew(c echo.Context) error {
	return render(c, "new.html", pongo2.Context{})
}

func handleDelete(c echo.Context) error {
	id, _ := strconv.Atoi(echo.Context.Param(c, "id"))
	movie := db.FetchMovieByID(int64(id))
	if !db.DeleteMovieByID(int64(id)) {
		os.Exit(2)
	}
	os.Remove("static/" + movie.Path)
	os.Remove("static/" + movie.Thumbnail)
	return c.Redirect(http.StatusFound, "/movies")
}

func handleUpload(c echo.Context) error {
	mov, _, merr := c.Request().FormFile("movie")
	img, _, ierr := c.Request().FormFile("thumbnail")
	if merr != nil || ierr != nil {
		return renderErr(c, merr)
	}
	defer mov.Close()
	defer img.Close()
	title := echo.Context.FormValue(c, "title")
	movOut, merr := os.Create("static/" + title + ".mp4")
	imgOut, ierr := os.Create("static/" + title + ".png")
	if merr != nil || ierr != nil {
		return renderErr(c, merr)
	}
	defer movOut.Close()
	defer imgOut.Close()
	io.Copy(movOut, mov)
	io.Copy(imgOut, img)

	if !db.InsertMoive(title, title+".mp4", title+".png") {
		os.Exit(2)
	}

	return c.Redirect(http.StatusFound, "/movies")
}

func handleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(echo.Context.Param(c, "id"))
	movie := db.FetchMovieByID(int64(id))
	return render(c, "edit.html", pongo2.Context{"movie": movie})
}

func render(c echo.Context, fileName string, ctx map[string]interface{}) error {
	tpl, err := templateSet.FromCache(fileName)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	pctx := pongo2.Context(ctx)
	pctx.Update(pongo2.Context{
		"c": c,
	})
	body, err := tpl.Execute(pctx)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, body)
}

func renderErr(c echo.Context, err error) error {
	return c.String(http.StatusInternalServerError, err.Error())
}
