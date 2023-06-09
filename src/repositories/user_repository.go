package repositories

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
	List(filter *models.Filter) ([]*models.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *models.User) error {
	query := `INSERT INTO users (
					username,
					password_hash,
					password_salt,
					deleted,
					deleted_at,
					created_at,
					updated_at
			) VALUES (
				$1,
				$2,
				$3,
				$4,
				$5,
				$6,
				$7
			)`

	_, err := r.db.Exec(query,
		user.Username,
		user.PasswordHash,
		user.PasswordSalt,
		false,
		nil,
		time.Time{},
		time.Time{},
	)

	if err != nil {
		return err
	}

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
	query := `UPDATE users SET deleted = $1, deleted_at = $2 WHERE id = $3`

	if _, err := r.db.Exec(query, true, time.Now(), 1); err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	return nil
}

func (r *UserRepositoryImpl) Find(id int) (*models.User, error) {
	query := `SELECT * FROM users WHERE id = $1`

	row := r.db.QueryRow(query, id)

	var user models.User

	if err := row.Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.PasswordSalt,
		&user.Deleted,
		&user.DeletedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}

	return &user, nil
}

func (r *UserRepositoryImpl) List(filter *models.Filter) ([]*models.User, error) {

	var where []string

	if filter != nil {
		if filter.Search != nil {
			where = append(where, fmt.Sprintf("username LIKE '%%%s%%'", *filter.Search))
		}

		if filter.From != nil {
			where = append(where, fmt.Sprintf("createdAt >= '%s'", *filter.From))
		}

		if filter.To != nil {
			where = append(where, fmt.Sprintf("createdAt <= '%s'", *filter.To))
		}

		if filter.Deleted != nil {
			where = append(where, fmt.Sprintf("deleted = %t", *filter.Deleted))
		}
	}

	query := `SELECT * FROM users`

	if len(where) > 0 {
		query += " WHERE " + where[0]
		for i := 1; i < len(where); i++ {
			query += " AND " + where[i]
		}
	}

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("failed to list users: %v", err)
	}

	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.PasswordHash,
			&user.PasswordSalt,
			&user.Deleted,
			&user.DeletedAt,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to list users: %v", err)
		}
		users = append(users, &user)
	}

	return users, nil
}
