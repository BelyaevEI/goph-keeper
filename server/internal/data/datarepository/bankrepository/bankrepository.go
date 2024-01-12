package bankrepository

import (
	"bytes"
	"context"
	"encoding/binary"
	"sync"

	"github.com/BelyaevEI/GophKeeper/server/internal/models/bankmodels"
	"github.com/BelyaevEI/GophKeeper/server/internal/storage/bankdb"
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
	var data bankmodels.Bankdata

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

func (bank *BankRepository) GetData(ctx context.Context, body []byte) ([]byte, error) {
	var (
		service bankmodels.Bankdata
		buf     bytes.Buffer
	)

	buffer := bytes.NewBuffer(body)

	// Deserializing binary data
	if err := binary.Read(buffer, binary.LittleEndian, &service); err != nil {
		return nil, err
	}

	// Locking for read data storage
	bank.mutex.Lock()

	defer bank.mutex.Unlock()

	data, err := bank.db.GetBankData(ctx, service)
	if err != nil {
		return nil, err
	}

	// Serializing data for return
	err = binary.Write(&buf, binary.LittleEndian, &data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}

func (pass *BankRepository) UpdateData(ctx context.Context, body []byte) error {

	var data bankmodels.Bankdata

	buffer := bytes.NewBuffer(body)

	// Deserializing binary data
	if err := binary.Read(buffer, binary.LittleEndian, &data); err != nil {
		return err
	}

	// Locking for read data storage
	pass.mutex.Lock()

	defer pass.mutex.Unlock()

	if err := pass.db.UpdateBankData(ctx, data); err != nil {
		return err
	}

	return nil
}

func (pass *BankRepository) DeleteData(ctx context.Context, body []byte) error {

	var data bankmodels.Bankdata

	buffer := bytes.NewBuffer(body)

	// Deserializing binary data
	if err := binary.Read(buffer, binary.LittleEndian, &data); err != nil {
		return err
	}

	// Locking for read data storage
	pass.mutex.Lock()

	defer pass.mutex.Unlock()

	if err := pass.db.DeleteBankData(ctx, data); err != nil {
		return err
	}

	return nil
}
