package passwordsdb

import (
	"context"
	"database/sql"

	"github.com/BelyaevEI/GophKeeper/server/internal/models/passwordsmodels"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Passwordsdb struct {
	db *sql.DB
}

type Store interface {
	SaveLR(ctx context.Context, data passwordsmodels.LRdata) error
	GetPassword(ctx context.Context, service passwordsmodels.LRdata) (passwordsmodels.LRdata, error)
	UpdatePassword(ctx context.Context, data passwordsmodels.LRdata) error
	DeletePassword(ctx context.Context, data passwordsmodels.LRdata) error
}

func NewConnect(dsn string) (*Passwordsdb, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS passwords
					 (userID int NOT NULL, 
					login text NOT NULL, 
					password text NOT NULL, 
					service text NOT NULL,
					note text NOT NULL)`)
	if err != nil {
		return nil, err
	}

	return &Passwordsdb{
		db: db,
	}, nil
}

func (passdb *Passwordsdb) SaveLR(ctx context.Context, data passwordsmodels.LRdata) error {

	_, err := passdb.db.ExecContext(ctx, "INSERT INTO passwords(userID, login, password, service, note)"+
		"values($1, $2, $3, $4)", data.UserID, data.Login, data.Password, data.Service, data.Note)
	return err
}

func (passdb *Passwordsdb) GetPassword(ctx context.Context, service passwordsmodels.LRdata) (passwordsmodels.LRdata, error) {

	var data passwordsmodels.LRdata

	row := passdb.db.QueryRowContext(ctx, "SELECT userID, login, password, service, note FROM passwords WHERE userID=$1 AND service=$2", service.UserID, service.Service)
	if err := row.Scan(&data); err != nil {
		return data, err
	}
	return data, nil
}

func (passdb *Passwordsdb) UpdatePassword(ctx context.Context, data passwordsmodels.LRdata) error {
	_, err := passdb.db.Exec("UPDATE passwords SET login = &1, password = &2, note = $3 WHERE userID = $4 AND service = $5",
		data.Login, data.Password, data.Note, data.UserID, data.Service)
	if err != nil {
		return err
	}
	return nil
}

func (passdb *Passwordsdb) DeletePassword(ctx context.Context, data passwordsmodels.LRdata) error {
	_, err := passdb.db.Exec("DELETE FROM passwords WHERE userID = $1 AND service = $2",
		data.UserID, data.Service)
	if err != nil {
		return err
	}
	return nil
}
