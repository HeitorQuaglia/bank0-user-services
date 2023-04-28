package userrepository

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
	"user-services/src/models"
)

type UserRepository interface {
	Create(*models.User) error
	Update(*models.User) error
	Delete(*models.User) error
	Find(int) (*models.User, error)
	List() ([]*models.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *models.User) error {
	return nil
}

func (r *UserRepositoryImpl) Update(user *models.User) error {
	if user.PasswordHash == nil {
		return errors.New("PasswordHash cannot be null")
	}

	saltBytes := make([]byte, 16)
	_, err := rand.Read(saltBytes)

	if err != nil {
		return fmt.Errorf("failed to generate salt: %v", err)
	}

	var salt = hex.EncodeToString(saltBytes)
	passwordHash := sha256.Sum256([]byte(*user.PasswordHash + salt))

	query := `UPDATE users SET password_hash = $1, password_salt = $2, updatedAt = $3 WHERE id = $4`

	if _, err = r.db.Exec(query, passwordHash, salt, time.Time{}, user.ID); err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}

func (r *UserRepositoryImpl) Delete(*models.User) error {
	return nil
}

func (r *UserRepositoryImpl) Find(id int) (*models.User, error) {
	return nil, nil
}

func (r *UserRepositoryImpl) List() ([]*models.User, error) {
	return nil, nil
}
