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
	//stmt, err := db.Prepare(`
	//	CREATE TABLE [music] (
	//	    id INTEGER NOT NULL PRIMARY KEY,
	//	    band TEXT NOT NULL,
	//	    song TEXT NOT NULL
	//	);
	//`)

	stmt, err := db.Prepare(
		"CREATE TABLE IF NOT EXISTS music (" +
			"id INTEGER PRIMARY KEY," +
			"`group` TEXT NOT NULL," +
			"song TEXT NOT NULL" +
			");")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, err
}

func (s *Storage) SaveMusic(group string, song string) (int64, error) {
	const op = "storage.sqlite.SaveMusic"

	fmt.Println("insert")
	stmt, err := s.db.Prepare("INSERT INTO music('group', 'song') VALUES (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	fmt.Println("insert exec")
	res, err := stmt.Exec(group, song)
	if err != nil {
		//if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
		//	return 0, fmt.Errorf("%s: %w", op, storage.ErrURLNotExist)
		//}
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
