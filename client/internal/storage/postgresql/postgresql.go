package postgresql

import (
	"context"
	"database/sql"
)

type Postgresql struct {
	db *sql.DB
}

func NewConnect(dsn string) (*Postgresql, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Create table for head data user
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS registation
					(login text NOT NULL, 
					password text NOT NULL, 
					secretkey text NOT NULL, 
					token text NOT NULL)`)
	if err != nil {
		return nil, err
	}

	return &Postgresql{
		db: db,
	}, nil

}

func (postgresql *Postgresql) CheckUniqueLogin(ctx context.Context, login string) error {
	_, err := postgresql.db.ExecContext(ctx, "SELECT login FROM registration WHERE login = $1", login)
	return err
}

func (postgresql *Postgresql) SaveDataNewUser(ctx context.Context, login, password, key string) error {
	_, err := postgresql.db.ExecContext(ctx, "INSERT INTO registration(login, password, secretkey)"+
		"values($1, $2, $3, $4)", login, password, key)
	return err
}

func (postgresql *Postgresql) GetSecretKey(ctx context.Context, userID uint32) (string, error) {
	var secretkey string

	row := postgresql.db.QueryRowContext(ctx, "SELECT secretkey FROM registration WHERE userID = $1", userID)
	if err := row.Scan(&secretkey); err != nil {
		return "", err
	}

	return secretkey, nil
}

func (postgresql *Postgresql) GetUserID(ctx context.Context, login string) (uint32, error) {
	var userID uint32

	row := postgresql.db.QueryRowContext(ctx, "SELECT userID FROM registration WHERE login = $1", login)
	if err := row.Scan(&userID); err != nil {
		return userID, err
	}

	return userID, nil
}

func (postgresql *Postgresql) GetPassword(ctx context.Context, login string) (string, error) {
	var password string

	row := postgresql.db.QueryRowContext(ctx, "SELECT password FROM registration WHERE login = $1", login)
	if err := row.Scan(&password); err != nil {
		return password, err
	}
	return password, nil
}
