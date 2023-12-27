package textrepository

import (
	"bytes"
	"context"
	"encoding/binary"
	"sync"

	"github.com/BelyaevEI/GophKeeper/client/internal/models"
	"github.com/BelyaevEI/GophKeeper/client/internal/storage/textsdb"
)

type TextRepository struct {
	mutex sync.RWMutex
	db    *textsdb.Textsdb
}

func New(db *textsdb.Textsdb) *TextRepository {
	return &TextRepository{
		db: db,
	}
}

func (text *TextRepository) SaveData(ctx context.Context, body []byte) error {

	var data models.Textsdata

	buffer := bytes.NewBuffer(body)

	// Deserializing binary data
	if err := binary.Read(buffer, binary.LittleEndian, &data); err != nil {
		return err
	}

	// Locking for safe data storage
	text.mutex.Lock()

	defer text.mutex.Unlock()

	// Save data in db
	err := text.db.SaveText(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
