package userdb

import (
	"context"
	"database/sql"
)

type UserDB struct {
	db *sql.DB
}

func NewConnect(dsn string) (*UserDB, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Create table for head data user
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS registation
					(login text NOT NULL, 
					password text NOT NULL, 
					secretkey text NOT NULL, 
					userID int NOT NULL`)
	if err != nil {
		return nil, err
	}

	return &UserDB{
		db: db,
	}, nil
}

func (userdb *UserDB) CheckUniqueLogin(ctx context.Context, login string) error {
	_, err := userdb.db.ExecContext(ctx, "SELECT login FROM registration WHERE login = $1", login)
	return err
}

func (userdb *UserDB) SaveDataNewUser(ctx context.Context, login, password, key string) error {
	_, err := userdb.db.ExecContext(ctx, "INSERT INTO registration(login, password, secretkey)"+
		"values($1, $2, $3, $4)", login, password, key)
	return err
}

func (userdb *UserDB) GetSecretKey(ctx context.Context, userID uint32) (string, error) {
	var secretkey string

	row := userdb.db.QueryRowContext(ctx, "SELECT secretkey FROM registration WHERE userID = $1", userID)
	if err := row.Scan(&secretkey); err != nil {
		return "", err
	}

	return secretkey, nil
}

func (userdb *UserDB) GetUserID(ctx context.Context, login string) (uint32, error) {
	var userID uint32

	row := userdb.db.QueryRowContext(ctx, "SELECT userID FROM registration WHERE login = $1", login)
	if err := row.Scan(&userID); err != nil {
		return userID, err
	}

	return userID, nil
}

func (userdb *UserDB) GetPassword(ctx context.Context, login string) (string, error) {
	var password string

	row := userdb.db.QueryRowContext(ctx, "SELECT password FROM registration WHERE login = $1", login)
	if err := row.Scan(&password); err != nil {
		return password, err
	}
	return password, nil
}
