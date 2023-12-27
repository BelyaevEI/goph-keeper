package storage

import (
	"github.com/BelyaevEI/GophKeeper/client/internal/storage/bankdb"
	"github.com/BelyaevEI/GophKeeper/client/internal/storage/bindb"
	"github.com/BelyaevEI/GophKeeper/client/internal/storage/passwordsdb"
	"github.com/BelyaevEI/GophKeeper/client/internal/storage/textsdb"
	"github.com/BelyaevEI/GophKeeper/client/internal/storage/userdb"
)

type DB struct {
	UserDB  *userdb.UserDB
	PassDB  *passwordsdb.Passwordsdb
	TextsDB *textsdb.Textsdb
	Bindb   *bindb.Bindb
	Bankdb  *bankdb.Bankdb
}

func Connect2DB(DSN string) (*DB, error) {

	// Connect to user data base
	userdb, err := userdb.NewConnect(DSN)
	if err != nil {
		return nil, err
	}

	// Connect to passwords data base
	passwordsdb, err := passwordsdb.NewConnect(DSN)
	if err != nil {
		return nil, err
	}

	// Connect to texts data base
	texts, err := textsdb.NewConnect(DSN)
	if err != nil {
		return nil, err
	}

	// Connect to texts data base
	bindb, err := bindb.NewConnect(DSN)
	if err != nil {
		return nil, err
	}

	// Connect to bank data base
	bankdb, err := bankdb.NewConnect(DSN)
	if err != nil {
		return nil, err
	}

	return &DB{
		UserDB:  userdb,
		PassDB:  passwordsdb,
		TextsDB: texts,
		Bindb:   bindb,
		Bankdb:  bankdb,
	}, nil
}
