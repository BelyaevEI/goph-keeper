package passwordrepository

import (
	"bytes"
	"context"
	"encoding/binary"
	"sync"

	"github.com/BelyaevEI/GophKeeper/client/internal/models/passwordsmodels"
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

	var data passwordsmodels.LRdata

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

func (pass *PasswordRepository) GetData(ctx context.Context, body []byte) ([]byte, error) {

	var (
		service passwordsmodels.LRdata
		buf     bytes.Buffer
	)

	buffer := bytes.NewBuffer(body)

	// Deserializing binary data
	if err := binary.Read(buffer, binary.LittleEndian, &service); err != nil {
		return nil, err
	}

	// Locking for read data storage
	pass.mutex.Lock()

	defer pass.mutex.Unlock()

	data, err := pass.db.GetPassword(ctx, service)
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

func (pass *PasswordRepository) UpdateData(ctx context.Context, body []byte) error {

	var data passwordsmodels.LRdata

	buffer := bytes.NewBuffer(body)

	// Deserializing binary data
	if err := binary.Read(buffer, binary.LittleEndian, &data); err != nil {
		return err
	}

	// Locking for read data storage
	pass.mutex.Lock()

	defer pass.mutex.Unlock()

	if err := pass.db.UpdatePassword(ctx, data); err != nil {
		return err
	}

	return nil
}
