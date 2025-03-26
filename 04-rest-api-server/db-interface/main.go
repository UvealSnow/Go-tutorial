package dbinterface

import (
	"database/sql"
	"os"

	"example.com/db-interface/repository"
	"github.com/go-sql-driver/mysql"
)

func ConnectToDB() (repositories *repository.RepositoryMap, err error) {
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_DATABASE"),
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	repositories = repository.NewRepositoryMap(db)

	return
}
