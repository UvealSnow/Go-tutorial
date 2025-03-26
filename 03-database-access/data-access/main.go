package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	db, err := connectToDB()
	if err != nil {
		log.Fatalf("Unable to connect: %q", err)
	}

	artist := "John Coltrane"
	albums, err := albumsByArtist(artist, db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums by artist %v found: %v\n", artist, albums)

	albumId := 2
	album, err := albumByID(int64(albumId), db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album by ID %v - found: %v\n", albumId, album)

	albId, err := addAlbum(
		Album{
			Title:  "Pneuma",
			Artist: "TOOL",
			Price:  69.96,
		},
		db,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of the added album: %v\n", albId)
}

func connectToDB() (db *sql.DB, err error) {
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_DATABASE"),
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return
}

func albumsByArtist(name string, db *sql.DB) (albums []Album, err error) {
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	// Defer schedules the execution of the function to just before this function returns.
	defer rows.Close()
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}

	// sql.Rows.Err() returns any error encountered during iteration.
	// This is the only way to find out if the actual query failed.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return
}

func albumByID(id int64, db *sql.DB) (album Album, err error) {
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("albumsById %d: Album not found", id)
		}
		return album, fmt.Errorf("albumsById %d: %v", id, err)
	}

	return album, nil
}

func addAlbum(alb Album, db *sql.DB) (id int64, err error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return
}
