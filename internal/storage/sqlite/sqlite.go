package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3" // init sqlite3 driver
	"github.com/xndrg/scheduly/internal/storage"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const fn = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS groups(
		id INTEGER PRIMARY KEY,
		chat_id UNSIGNED BIG INT NOT NULL UNIQUE,
		study_group TEXT NOT NULL);
	CREATE INDEX IF NOT EXISTS idx_chat_id ON groups(chat_id);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveGroup(chatID int64, group string) (int64, error) {
	const fn = "storage.sqlite.SaveGroup"

	stmt, err := s.db.Prepare("INSERT INTO groups(chat_id, study_group) VALUES(?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: prepate statement: %w", fn, err)
	}

	res, err := stmt.Exec(chatID, group)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("%s: %w", fn, storage.ErrGroupAlreadyEsists)
		}

		return 0, fmt.Errorf("%s: execute statement: %w", fn, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", fn, err)
	}

	return id, err
}

func (s *Storage) GetGroup(chatID int64) (string, error) {
	const fn = "storage.sqlite.GetGroup"

	stmt, err := s.db.Prepare("SELECT study_group FROM groups WHERE chat_id = ?")
	if err != nil {
		return "", fmt.Errorf("%s: prepare statement: %w", fn, err)
	}

	var resGroup string

	err = stmt.QueryRow(chatID).Scan(&resGroup)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", storage.ErrGroupNotFound
		}

		return "", fmt.Errorf("%s: execute statement: %w", fn, err)
	}

	return resGroup, nil
}

// TODO: implement methods
// func (s *Storage) DeleteGroup(chatID int64) error
// func (s *Storage) ChangeGroup(chatID int64, newGroup string) error
