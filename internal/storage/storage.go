package storage

import "errors"

var (
	ErrGroupNotFound      = errors.New("group not found")
	ErrGroupAlreadyEsists = errors.New("group already exists")
)

type Storage interface {
	SaveGroup(chatID int64, group string) (int64, error)
	GetGroup(chatID int64) (string, error)
}
