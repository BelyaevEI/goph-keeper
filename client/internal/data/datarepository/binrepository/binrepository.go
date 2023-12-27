package binrepository

import (
	"bytes"
	"context"
	"encoding/binary"
	"sync"

	"github.com/BelyaevEI/GophKeeper/client/internal/models"
	"github.com/BelyaevEI/GophKeeper/client/internal/storage/bindb"
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

	var data models.Binarydata

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
