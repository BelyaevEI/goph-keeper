package passwordrepository

import (
	"bytes"
	"context"
	"encoding/binary"
	"sync"

	"github.com/BelyaevEI/GophKeeper/client/internal/models"
	"github.com/BelyaevEI/GophKeeper/client/internal/storage/passwordsdb"
)

type PasswordRepository struct {
	mutex sync.RWMutex
	db    *passwordsdb.Passwordsdb
}

func New(db *passwordsdb.Passwordsdb) *PasswordRepository {
	return &PasswordRepository{
		db: db,
	}
}

func (pass *PasswordRepository) SaveData(ctx context.Context, body []byte) error {

	var data models.LRdata

	buffer := bytes.NewBuffer(body)

	// Deserializing binary data
	if err := binary.Read(buffer, binary.LittleEndian, &data); err != nil {
		return err
	}

	// Locking for safe data storage
	pass.mutex.Lock()

	defer pass.mutex.Unlock()

	// Save data in db
	err := pass.db.SaveLR(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
