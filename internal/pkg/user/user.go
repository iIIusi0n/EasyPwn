package user

import (
	"context"
	"database/sql"
	"easypwn/internal/pkg/util"
	"time"
)

type User struct {
	ID        string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(ctx context.Context, db *sql.DB, email, password string) (*User, error) {
	passwordHash := util.HashPassword(password)

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO user (id, email, password_hash) VALUES (UUID_TO_BIN(UUID()), ?, ?)", email, passwordHash)
	if err != nil {
		return nil, err
	}

	var userId string
	err = tx.QueryRow("SELECT BIN_TO_UUID(id) FROM user WHERE email = ?", email).Scan(&userId)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(`
		INSERT INTO user_license (id, user_id, license_type_id) 
		SELECT UUID_TO_BIN(UUID()), UUID_TO_BIN(?), id FROM user_license_type WHERE name = 'free'
	`, userId)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &User{
		ID:        userId,
		Email:     email,
		Password:  passwordHash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func GetUser(ctx context.Context, db *sql.DB, id string) (*User, error) {
	var createdAt, updatedAt string
	row := db.QueryRow("SELECT BIN_TO_UUID(id), email, password_hash, DATE_FORMAT(created_at, '%Y-%m-%dT%H:%i:%sZ'), DATE_FORMAT(updated_at, '%Y-%m-%dT%H:%i:%sZ') FROM user WHERE id = UUID_TO_BIN(?)", id)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	user.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return nil, err
	}
	user.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(ctx context.Context, db *sql.DB, email string) (*User, error) {
	var createdAt, updatedAt string
	row := db.QueryRow("SELECT BIN_TO_UUID(id), email, password_hash, DATE_FORMAT(created_at, '%Y-%m-%dT%H:%i:%sZ'), DATE_FORMAT(updated_at, '%Y-%m-%dT%H:%i:%sZ') FROM user WHERE email = ?", email)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	user.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return nil, err
	}
	user.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Delete(ctx context.Context, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM user WHERE id = UUID_TO_BIN(?)", u.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (u *User) GetLicense(ctx context.Context, db *sql.DB) (string, error) {
	var licenseType string
	err := db.QueryRow(`
		SELECT lt.name 
		FROM user_license ul
		JOIN user_license_type lt ON lt.id = ul.license_type_id 
		WHERE ul.user_id = UUID_TO_BIN(?)`, u.ID).Scan(&licenseType)
	if err != nil {
		return "", err
	}

	return licenseType, nil
}

func (c *User) UpdateLicense(ctx context.Context, db *sql.DB, licenseType string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		UPDATE user_license 
		SET license_type_id = (SELECT id FROM user_license_type WHERE name = ?)
		WHERE user_id = UUID_TO_BIN(?)
	`, licenseType, c.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (c *User) UpdatePassword(ctx context.Context, db *sql.DB, password string) error {
	passwordHash := util.HashPassword(password)

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE user SET password_hash = ? WHERE id = UUID_TO_BIN(?)", passwordHash, c.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
