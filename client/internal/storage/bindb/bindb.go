package bindb

import (
	"context"
	"database/sql"

	"github.com/BelyaevEI/GophKeeper/client/internal/models"
)

type Bindb struct {
	db *sql.DB
}

func NewConnect(dsn string) (*Bindb, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS binary
					(userID int NOT NULL, 
					bin text NOT NULL, 
					service text NOT NULL,
					note text NOT NULL)`)
	if err != nil {
		return nil, err
	}

	return &Bindb{db: db}, nil
}

func (bindb *Bindb) SaveBin(ctx context.Context, data models.Binarydata) error {

	_, err := bindb.db.ExecContext(ctx, "INSERT INTO texts(userID, bin, service, note)"+
		"values($1, $2, $3, $4)", data.UserID, data.Bin, data.Service, data.Note)
	return err
}
