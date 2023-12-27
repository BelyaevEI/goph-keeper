package passwordsdb

import (
	"context"
	"database/sql"

	"github.com/BelyaevEI/GophKeeper/client/internal/models"
)

type Passwordsdb struct {
	db *sql.DB
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

func (passdb *Passwordsdb) SaveLR(ctx context.Context, data models.LRdata) error {

	_, err := passdb.db.ExecContext(ctx, "INSERT INTO passwords(userID, login, password, service, note)"+
		"values($1, $2, $3, $4)", data.UserID, data.Login, data.Password, data.Service, data.Note)
	return err
}
