package album

import (
	"database/sql"
	"fmt"
)

// @TODO: Played with the idea of creating a basic migration system.
// var TableDefinition = repository.TableDefinition{
// 	Name: "album",
// 	Columns: repository.FieldMap{
// 		"id": {
// 			Type:     "INT AUTOINCREMENT",
// 			Nullable: false,
// 			Constraints: []repository.Constraint{
// 				{Type: "PRIMARY_KEY", Value: ""},
// 				{Type: "AUTO_INCREMENT", Value: ""},
// 			},
// 		},
// 		"title": {
// 			Type:     "VARCHAR(255)",
// 			Nullable: false,
// 		},
// 		"artist": {
// 			Type:     "VARCHAR(255)",
// 			Nullable: false,
// 		},
// 		"price": {
// 			Type:     "DECIMALS(5,2)",
// 			Nullable: false,
// 			Constraints: []repository.Constraint{
// 				{Type: "CHECK", Value: "price > 0"},
// 			},
// 		},
// 	},
// }

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

func (album Album) GetId() int64 {
	return album.ID
}

func (album Album) SetId(id int64) {
	album.ID = id
}

type AlbumRepository struct {
	db *sql.DB
}

func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	repo := new(AlbumRepository)
	repo.db = db
	return repo
}

func (repo AlbumRepository) FindById(id int64) (album Album, err error) {
	row := repo.db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("AlbumsRepository.FindById: %d - Not found", id)
		}
		return album, fmt.Errorf("AlbumsRepository.FindById: %d - Error Scanning", id)
	}
	return
}

func (repo AlbumRepository) FindAll() (albums []Album, err error) {
	rows, err := repo.db.Query("SELECT * FROM album")
	if err != nil {
		return nil, fmt.Errorf("AlbumRepository.FindAll: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("AlbumRepository.FindAll: %v", err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("AlbumRepository.FindAll: %v", err)
	}

	return
}

func (repo AlbumRepository) Insert(alb *Album) (albumId int64, err error) {
	result, err := repo.db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", &alb.Title, &alb.Artist, &alb.Price)
	if err != nil {
		return 0, fmt.Errorf("AlbumRepository.Insert: %v", err)
	}

	albumId, err = result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AlbumRepository.Insert: %v", err)
	}

	return
}

func (repo AlbumRepository) UpdateIfExists(id int64, alb *Album) (album Album, err error) {
	result, err := repo.db.Exec("UPDATE album SET title = ?, artist = ?, price = ? WHERE ID = ?", &alb.Title, &alb.Artist, &alb.Price)
	if err != nil {
		return album, fmt.Errorf("AlbumRepository.UpdateIfExists: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return album, fmt.Errorf("AlbumRepository.UpdateIfExists: %v", err)
	}

	if rowsAffected == 0 || rowsAffected > 1 {
		return album, fmt.Errorf("AlbumRepository.UpdateIfExists: Rows affected %v", rowsAffected)
	}

	album, err = repo.FindById(id)
	if err != nil {
		return album, fmt.Errorf("AlbumRepository.UpdateIfExists: %v", err)
	}

	return
}

func (repo AlbumRepository) DeleteIfExists(id int64) (album Album, err error) {
	album, err = repo.FindById(id)
	if err != nil {
		return album, fmt.Errorf("AlbumRepository.DeleteIfExists: %v", err)
	}
	album.SetId(0)

	result, err := repo.db.Exec("DELETE FROM album WHERE id = ?", id)
	if err != nil {
		return album, fmt.Errorf("AlbumRepository.DeleteIfExists: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return album, fmt.Errorf("AlbumRepository.DeleteIfExists: %v", err)
	}

	if rowsAffected == 0 {
		return album, fmt.Errorf("AlbumRepository.DeleteIfExists: album with ID %v not found", id)
	}

	return
}
