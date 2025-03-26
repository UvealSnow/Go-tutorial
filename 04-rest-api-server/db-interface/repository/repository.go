package repository

import (
	"database/sql"

	"example.com/db-interface/album"
)

type Model interface {
	GetId() int64
	SetId(id int64)
}

type Constraint struct {
	Type  string
	Value string
}

// type Field struct {
// 	Type        string
// 	Nullable    bool
// 	Constraints []Constraint
// }

// type FieldMap map[string]Field

// type TableDefinition struct {
// 	Name    string
// 	Columns FieldMap
// }

type Repository[T Model] interface {
	FindById(id int64) (T, error)
	FindAll() ([]T, error)
	Insert(model *T) (int64, error)
	UpdateIfExists(id int64, model *T) (T, error)
	DeleteIfExists(id int64) (T, error)
}

type RepositoryMap struct {
	Albumrepository *album.AlbumRepository
}

func NewRepositoryMap(db *sql.DB) *RepositoryMap {
	repositoryMap := new(RepositoryMap)
	repositoryMap.Albumrepository = album.NewAlbumRepository(db)
	return repositoryMap
}
