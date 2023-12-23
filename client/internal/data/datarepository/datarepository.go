package datarepository

import "github.com/BelyaevEI/GophKeeper/client/internal/storage/postgresql"

type DataRepository struct {
	db *postgresql.Postgresql
}

func New(db *postgresql.Postgresql) *DataRepository {
	return &DataRepository{
		db: db,
	}
}
