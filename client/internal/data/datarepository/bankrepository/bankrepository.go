package bankrepository

import (
	"bytes"
	"context"
	"encoding/binary"
	"sync"

	"github.com/BelyaevEI/GophKeeper/client/internal/models"
	"github.com/BelyaevEI/GophKeeper/client/internal/storage/bankdb"
)

type BankRepository struct {
	db    *bankdb.Bankdb
	mutex sync.RWMutex
}

func New(db *bankdb.Bankdb) *BankRepository {
	return &BankRepository{
		db: db,
	}
}

func (bank *BankRepository) SaveData(ctx context.Context, body []byte) error {
	var data models.Bankdata

	buffer := bytes.NewBuffer(body)

	// Deserializing binary data
	if err := binary.Read(buffer, binary.LittleEndian, &data); err != nil {
		return err
	}

	// Locking for safe data storage
	bank.mutex.Lock()

	defer bank.mutex.Unlock()

	// Save data in db
	err := bank.db.SaveBank(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
