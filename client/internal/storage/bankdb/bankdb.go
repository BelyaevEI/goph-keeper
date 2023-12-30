package bankdb

import (
	"context"
	"database/sql"

	"github.com/BelyaevEI/GophKeeper/client/internal/models/bankmodels"
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

func (bank *Bankdb) SaveBank(ctx context.Context, data bankmodels.Bankdata) error {
	_, err := bank.db.ExecContext(ctx, "INSERT INTO bank(userID, fullname, number, date, cvc, bankname, note)"+
		"values($1, $2, $3, $4, $5, $6, $7)", data.UserID, data.Fullname, data.Number, data.Date, data.Cvc, data.Bankname, data.Note)
	return err
}

func (bank *Bankdb) GetBankData(ctx context.Context, service bankmodels.Bankdata) (bankmodels.Bankdata, error) {

	var data bankmodels.Bankdata

	row := bank.db.QueryRowContext(ctx, "SELECT userID, fullname, number, date, cvc, bankname, note FROM bank WHERE userID=$1 AND bankname=$2",
		service.UserID, service.Bankname)
	if err := row.Scan(&data); err != nil {
		return data, err
	}
	return data, nil
}

func (bank *Bankdb) UpdateBankData(ctx context.Context, data bankmodels.Bankdata) error {
	_, err := bank.db.Exec("UPDATE bank SET fullname = &1, number = &2, date = $3, cvc = $4, note = $5 WHERE userID = $6 AND bankname = $7",
		data.Fullname, data.Number, data.Date, data.Cvc, data.Note, data.UserID, data.Bankname)
	if err != nil {
		return err
	}
	return nil
}
