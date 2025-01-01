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
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(ctx context.Context, db *sql.DB, email, username, password string) (*User, error) {
	passwordHash := util.HashPassword(password)

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var userId string
	result := tx.QueryRow("INSERT INTO user (email, username, password_hash) VALUES (?, ?, ?) RETURNING id", email, username, passwordHash)
	err = result.Scan(&userId)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(`
		INSERT INTO user_license (user_id, license_type_id) 
		SELECT ?, id FROM user_license_type WHERE name = 'free'
	`, userId)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &User{ID: userId}, nil
}

func GetUser(ctx context.Context, db *sql.DB, id string) (*User, error) {
	row := db.QueryRow("SELECT id, email, username, password_hash, created_at, updated_at FROM user WHERE id = ?", id)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(ctx context.Context, db *sql.DB, email string) (*User, error) {
	row := db.QueryRow("SELECT id, email, username, password_hash, created_at, updated_at FROM user WHERE email = ?", email)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
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

	_, err = tx.Exec("DELETE FROM user WHERE id = ?", u.ID)
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
		WHERE ul.user_id = ?`, u.ID).Scan(&licenseType)
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
		WHERE user_id = ?
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

	_, err = tx.Exec("UPDATE user SET password_hash = ? WHERE id = ?", passwordHash, c.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
