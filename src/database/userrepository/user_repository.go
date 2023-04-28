package userrepository

import (
	"database/sql"
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
