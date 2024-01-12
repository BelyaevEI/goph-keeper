package binrepository

import (
	"bytes"
	"context"
	"encoding/binary"
	"sync"

	"github.com/BelyaevEI/GophKeeper/server/internal/models/binarymodels"
	"github.com/BelyaevEI/GophKeeper/server/internal/storage/bindb"
)

type BinRepository struct {
	mutex sync.RWMutex
	db    *bindb.Bindb
}

func New(db *bindb.Bindb) *BinRepository {
	return &BinRepository{
		db: db,
	}
}

func (br *BinRepository) SaveData(ctx context.Context, body []byte) error {

	var data binarymodels.Binarydata

	buffer := bytes.NewBuffer(body)

	// Deserializing binary data
	if err := binary.Read(buffer, binary.LittleEndian, &data); err != nil {
		return err
	}

	// Locking for safe data storage
	br.mutex.Lock()

	defer br.mutex.Unlock()

	// Save data in db
	err := br.db.SaveBin(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (bin *BinRepository) GetData(ctx context.Context, body []byte) ([]byte, error) {

	var (
		service binarymodels.Binarydata
		buf     bytes.Buffer
	)

	buffer := bytes.NewBuffer(body)

	// Deserializing binary data
	if err := binary.Read(buffer, binary.LittleEndian, &service); err != nil {
		return nil, err
	}

	// Locking for read data storage
	bin.mutex.Lock()

	defer bin.mutex.Unlock()

	data, err := bin.db.GetBinary(ctx, service)
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

func (pass *BinRepository) UpdateData(ctx context.Context, body []byte) error {

	var data binarymodels.Binarydata

	buffer := bytes.NewBuffer(body)

	// Deserializing binary data
	if err := binary.Read(buffer, binary.LittleEndian, &data); err != nil {
		return err
	}

	// Locking for read data storage
	pass.mutex.Lock()

	defer pass.mutex.Unlock()

	if err := pass.db.UpdateBinary(ctx, data); err != nil {
		return err
	}

	return nil
}

func (pass *BinRepository) DeleteData(ctx context.Context, body []byte) error {

	var data binarymodels.Binarydata

	buffer := bytes.NewBuffer(body)

	// Deserializing binary data
	if err := binary.Read(buffer, binary.LittleEndian, &data); err != nil {
		return err
	}

	// Locking for read data storage
	pass.mutex.Lock()

	defer pass.mutex.Unlock()

	if err := pass.db.DeleteBinary(ctx, data); err != nil {
		return err
	}

	return nil

}
