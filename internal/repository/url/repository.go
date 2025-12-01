package url

import (
	"errors"

	"github.com/wb-go/wbf/dbpg"
)

var ErrFullURLNotFound = errors.New("full url not found")

type RepositoryURL struct {
	db *dbpg.DB
}

func NewRepository(db *dbpg.DB) *RepositoryURL {
	return &RepositoryURL{
		db: db,
	}
}
