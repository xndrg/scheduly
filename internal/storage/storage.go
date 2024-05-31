package storage

import "errors"

var (
	ErrGroupNotFound      = errors.New("group not found")
	ErrGroupAlreadyEsists = errors.New("group already exists")
)
