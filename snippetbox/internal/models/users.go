package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
	VALUES (?, ?, ?, CURRENT_TIMESTAMP)`

	_, err = m.DB.Exec(stmt, name, email, hashedPassword)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			// better way?
			if sqliteErr.Code == sqlite3.ErrConstraint && strings.Contains(sqliteErr.Error(), "UNIQUE constraint failed: users.email") {
				return ErrDuplicateEmail
			}
		}
		return err
	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	var (
		id             int
		hashedPassword []byte
	)

	stmt := `SELECT id, hashed_password FROM users WHERE email = ?`

	err := m.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		}
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		}
		return 0, err
	}

	return id, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	var exists bool

	stmt := `SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)`

	err := m.DB.QueryRow(stmt, id).Scan(&exists)
	return exists, err
}
