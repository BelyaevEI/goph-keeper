package textsdb

import (
	"context"
	"database/sql"

	"github.com/BelyaevEI/GophKeeper/server/internal/models/textsmodels"
	_ "github.com/jackc/pgx/v5/stdlib"
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

func (textsdb *Textsdb) SaveText(ctx context.Context, data textsmodels.Textsdata) error {

	_, err := textsdb.db.ExecContext(ctx, "INSERT INTO texts(userID, text, service, note)"+
		"values($1, $2, $3, $4)", data.UserID, data.Text, data.Service, data.Note)
	return err
}

func (textsdb *Textsdb) GetTexts(ctx context.Context, service textsmodels.Textsdata) (textsmodels.Textsdata, error) {

	var data textsmodels.Textsdata

	row := textsdb.db.QueryRowContext(ctx, "SELECT userID, text, service, note FROM texts WHERE userID=$1 AND service=$2", service.UserID, service.Service)
	if err := row.Scan(&data); err != nil {
		return data, err
	}
	return data, nil

}

func (passdb *Textsdb) UpdateText(ctx context.Context, data textsmodels.Textsdata) error {
	_, err := passdb.db.Exec("UPDATE texts SET text = &1, note = $2 WHERE userID = $3 AND service = $4",
		data.Text, data.Note, data.Note, data.UserID, data.Service)
	if err != nil {
		return err
	}
	return nil
}

func (passdb *Textsdb) DeleteText(ctx context.Context, data textsmodels.Textsdata) error {
	_, err := passdb.db.Exec("DELETE FROM texts WHERE userID = $1 AND service = $2",
		data.UserID, data.Service)
	if err != nil {
		return err
	}
	return nil
}
