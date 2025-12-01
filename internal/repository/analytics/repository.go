package analytics

import "github.com/wb-go/wbf/dbpg"

type RepositoryAnalytics struct {
	db *dbpg.DB
}

func NewRepository(db *dbpg.DB) *RepositoryAnalytics {
	return &RepositoryAnalytics{
		db: db,
	}
}
