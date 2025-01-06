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
	FileName  string
	OsID      string
	PluginID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProject(ctx context.Context, db *sql.DB, name, userID, filePath, fileName, osID, pluginID string) (*Project, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var projectID string
	err = tx.QueryRow("SELECT UUID() INTO @uuid").Scan(&projectID)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec("INSERT INTO project (id, name, user_id, file_path, file_name, os_id, plugin_id) VALUES (UUID_TO_BIN(@uuid), ?, ?, ?, ?, ?, ?)", name, userID, filePath, fileName, osID, pluginID)
	if err != nil {
		return nil, err
	}

	err = tx.QueryRow("SELECT @uuid").Scan(&projectID)
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
		FileName:  fileName,
		OsID:      osID,
		PluginID:  pluginID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func GetProject(ctx context.Context, db *sql.DB, id string) (*Project, error) {
	project := &Project{}
	err := db.QueryRow("SELECT * FROM project WHERE id = UUID_TO_BIN(?)", id).Scan(
		&project.ID,
		&project.Name,
		&project.UserID,
		&project.FilePath,
		&project.FileName,
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

	_, err = tx.Exec("DELETE FROM project WHERE id = UUID_TO_BIN(?)", p.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
