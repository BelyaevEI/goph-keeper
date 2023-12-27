package bankdb

import (
	"context"
	"database/sql"

	"github.com/BelyaevEI/GophKeeper/client/internal/models"
)

type Bankdb struct {
	db *sql.DB
}

func NewConnect(dsn string) (*Bankdb, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS bank
					 (userID int NOT NULL, 
					fullname text NOT NULL, 
					number text NOT NULL, 
					date text NOT NULL,
					cvc int NOT NULL,
					bankname text NOT NULL,
					note text NOT NULL)`)
	if err != nil {
		return nil, err
	}

	return &Bankdb{
		db: db,
	}, nil

}

func (bank *Bankdb) SaveBank(ctx context.Context, data models.Bankdata) error {
	_, err := bank.db.ExecContext(ctx, "INSERT INTO bank(userID, fullname, number, date, cvc, bankname, note)"+
		"values($1, $2, $3, $4, $5, $6, $7)", data.UserID, data.Fullname, data.Number, data.Date, data.Cvc, data.Bankname, data.Note)
	return err
}
