package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"
	//storagePath
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	//defer db.Close()
	stmt, err := db.Prepare(`
		CREATE TABLE [music] (
		    id INTEGER NOT NULL PRIMARY KEY,
		    band TEXT NOT NULL,
		    song TEXT NOT NULL
		);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	//stmt, err := db.Prepare(
	//	"CREATE TABLE IF NOT EXISTS music (" +
	//		"id INTEGER PRIMARY KEY," +
	//		"`group` TEXT NOT NULL," +
	//		"song TEXT NOT NULL" +
	//		");")
	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, err
}
