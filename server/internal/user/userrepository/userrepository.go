package userrepository

import (
	"context"
	"database/sql"
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/BelyaevEI/GophKeeper/server/internal/models"
	"github.com/BelyaevEI/GophKeeper/server/internal/storage/userdb"
	"golang.org/x/crypto/bcrypt"
)

type Store interface {
	GenerateUniqueUserID() (uint32, error)
	GenerateRandomString(length int) string
	SaveDataNewUser(ctx context.Context, login, password, key string, userID uint32) error
	CheckUniqueLogin(ctx context.Context, login string) bool
}

type UserRepository struct {
	db userdb.Store
}

func New(db userdb.Store) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Checking there is a user with this username
func (repository *UserRepository) CheckUniqueLogin(ctx context.Context, login string) bool {
	err := repository.db.CheckUniqueLogin(ctx, login)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		fmt.Print(err)
		return false
	}
	return true
}

func (repository *UserRepository) GenerateUniqueUserID() (uint32, error) {

	time := time.Now().UnixNano()
	randomByte := make([]byte, 4)

	_, err := rand.Read(randomByte)
	if err != nil {
		return 0, err
	}

	return uint32(time) + binary.BigEndian.Uint32(randomByte), nil
}

func (repository *UserRepository) GenerateRandomString(length int) string {

	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = models.CharSet[rand.Intn(len(models.CharSet))]
	}

	return string(result)
}

// saving registered user data
func (repository *UserRepository) SaveDataNewUser(ctx context.Context, login, password, key string, userID uint32) error {

	// Generate hash password with salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = repository.db.SaveDataNewUser(ctx, login, string(hashedPassword), key, userID)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepository) GetSecretKey(ctx context.Context, userID uint32) (string, error) {

	// we will get the user's password from the database
	secretkey, err := repository.db.GetSecretKey(ctx, userID)
	if err != nil {
		return "", err
	}

	return secretkey, nil
}

func (repository *UserRepository) GetUserID(ctx context.Context, login string) (uint32, error) {

	// we will get the user's password from the database
	userID, err := repository.db.GetUserID(ctx, login)
	if err != nil {
		return userID, err
	}
	return userID, nil
}

func (repository *UserRepository) VerifyingPassword(ctx context.Context, login, password string) (bool, error) {

	hashPassword, err := repository.db.GetPassword(ctx, login)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}
