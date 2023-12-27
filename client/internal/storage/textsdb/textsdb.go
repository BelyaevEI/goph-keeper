package textsdb

import (
	"context"
	"database/sql"

	"github.com/BelyaevEI/GophKeeper/client/internal/models"
)

type Textsdb struct {
	db *sql.DB
}

func NewConnect(dsn string) (*Textsdb, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS texts
					(userID int NOT NULL, 
					text text NOT NULL, 
					service text NOT NULL,
					note text NOT NULL)`)

	if err != nil {
		return nil, err
	}

	return &Textsdb{
		db: db,
	}, nil
}

func (textsdb *Textsdb) SaveText(ctx context.Context, data models.Textsdata) error {

	_, err := textsdb.db.ExecContext(ctx, "INSERT INTO texts(userID, text, service, note)"+
		"values($1, $2, $3, $4)", data.UserID, data.Text, data.Service, data.Note)
	return err
}
