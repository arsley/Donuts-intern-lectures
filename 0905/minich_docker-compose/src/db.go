package main

import (
	"fmt"
	"log"
	"os"
)

type Movie struct {
	ID        int64  `db:"id"`
	Title     string `db:"title"`
	Path      string `db:"path"`
	Thumbnail string `db:"thumbnail"`
}

// FetchMovies : fetch all movie and then print its records
func (db DB) FetchMovies() []*Movie {
	var movies []*Movie
	q := " SELECT * FROM `movie` "
	if err := db.Select(&movies, q); err != nil {
		log.Fatalf("FetchMovies Error: %v\n", err)
		os.Exit(2)
	}

	return movies

	// fmt.Printf("Number of Moives is %d\n", len(movies))
	// for _, m := range movies {
	// 	fmt.Printf("MoiveID:%d, Title:%s\n", m.ID, m.Title)
	// }
	// return
}

// FetchMovieByID : fetch movie by id and then print its record
func (db DB) FetchMovieByID(id int64) Movie {
	var movie Movie
	q := " SELECT * FROM `movie` WHERE `id` = ? "
	if err := db.Get(&movie, q, id); err != nil {
		log.Fatalf("FetchMovieByID Error: %v\n", err)
		os.Exit(2)
	}
	// fmt.Printf("MoiveID:%d, Title:%s\n", movie.ID, movie.Title)
	return movie
}

func (db DB) DeleteMovieByID(id int64) bool {
	q := "DELETE FROM `movie` WHERE `id` = ?"
	res, err := db.Exec(q, id)
	if err != nil {
		log.Fatalf("DeleteMovieByID Error: %v\n", err)
		return false
	}
	if rowsAffected, err := res.RowsAffected(); err != nil {
		log.Fatalf("DeleteMovieByID Error: %v\n", err)
		return false
	} else if rowsAffected > 0 {
		fmt.Printf("Successfully deleted. ID:%d\n", id)
		return true
	} else {
		fmt.Printf("ID:%d doesn't exist\n", id)
		return false
	}
}

func (db DB) UpdateMovieTitleByID(id int64, title string) {
	q := "UPDATE `movie` SET `title` = ? WHERE `id` = ?"
	res, err := db.Exec(q, title, id)
	if err != nil {
		log.Fatalf("DeleteMovieByID Error: %v\n", err)
		return
	}
	if rowsAffected, err := res.RowsAffected(); err != nil {
		log.Fatalf("UpdateMovieTitleByID Error: %v\n", err)
	} else if rowsAffected > 0 {
		fmt.Printf("Successfully Updated. ID:%d\n", id)
	} else {
		fmt.Printf("ID:%d doesn't exist or title is same.\n", id)
	}
	return
}

func (db DB) InsertMoive(title string, path string, thumbnail string) bool {
	q := "INSERT INTO `movie` (`title`, `path`, `thumbnail`) VALUES (?, ?, ?) "
	res, err := db.Exec(q, title, path, thumbnail)
	if err != nil {
		log.Fatalf("InsertMoive Error: %v\n", err)
		return false
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("InsertMoive Error: %v\n", err)
		return false
	}
	fmt.Printf("Successfully Inserted. ID is %d\n", lastInsertedID)
	return true
}
