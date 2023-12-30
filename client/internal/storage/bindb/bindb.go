package bindb

import (
	"context"
	"database/sql"

	"github.com/BelyaevEI/GophKeeper/client/internal/models/binarymodels"
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

func (bindb *Bindb) SaveBin(ctx context.Context, data binarymodels.Binarydata) error {

	_, err := bindb.db.ExecContext(ctx, "INSERT INTO texts(userID, bin, service, note)"+
		"values($1, $2, $3, $4)", data.UserID, data.Bin, data.Service, data.Note)
	return err
}

func (bindb *Bindb) GetBinary(ctx context.Context, service binarymodels.Binarydata) (binarymodels.Binarydata, error) {

	var data binarymodels.Binarydata

	row := bindb.db.QueryRowContext(ctx, "SELECT userID, bin, service, note FROM texts WHERE userID=$1 AND service=$2", service.UserID, service.Service)
	if err := row.Scan(&data); err != nil {
		return data, err
	}
	return data, nil
}
