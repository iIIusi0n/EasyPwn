package project

import (
	"context"
	"database/sql"
	"time"
)

type Project struct {
	ID        string
	Name      string
	UserID    string
	FilePath  string
	OsID      string
	PluginID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProject(ctx context.Context, db *sql.DB, name, userID, filePath, osID, pluginID string) (*Project, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var projectID string
	result := tx.QueryRow("INSERT INTO project (name, user_id, file_path, os_id, plugin_id) VALUES (?, ?, ?, ?, ?) RETURNING id", name, userID, filePath, osID, pluginID)
	err = result.Scan(&projectID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &Project{
		ID:        projectID,
		Name:      name,
		UserID:    userID,
		FilePath:  filePath,
		OsID:      osID,
		PluginID:  pluginID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func GetProject(ctx context.Context, db *sql.DB, id string) (*Project, error) {
	project := &Project{}
	err := db.QueryRow("SELECT * FROM project WHERE id = $1", id).Scan(
		&project.ID,
		&project.Name,
		&project.UserID,
		&project.FilePath,
		&project.OsID,
		&project.PluginID,
		&project.CreatedAt,
		&project.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (p *Project) Delete(ctx context.Context, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM project WHERE id = ?", p.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
