package textrepository

import (
	"bytes"
	"context"
	"encoding/binary"
	"sync"

	"github.com/BelyaevEI/GophKeeper/client/internal/models/textsmodels"
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

	var data textsmodels.Textsdata

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

func (text *TextRepository) GetData(ctx context.Context, body []byte) ([]byte, error) {

	var (
		service textsmodels.Textsdata
		buf     bytes.Buffer
	)

	buffer := bytes.NewBuffer(body)

	// Deserializing binary data
	if err := binary.Read(buffer, binary.LittleEndian, &service); err != nil {
		return nil, err
	}

	// Locking for read data storage
	text.mutex.Lock()

	defer text.mutex.Unlock()

	data, err := text.db.GetTexts(ctx, service)
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
