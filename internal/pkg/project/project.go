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

	_, err = tx.Exec("SELECT UUID() INTO @uuid")
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec("INSERT INTO project (id, name, user_id, file_path, file_name, os_id, plugin_id) VALUES (UUID_TO_BIN(@uuid), ?, UUID_TO_BIN(?), ?, ?, UUID_TO_BIN(?), UUID_TO_BIN(?)", name, userID, filePath, fileName, osID, pluginID)
	if err != nil {
		return nil, err
	}

	var projectID string
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
	var createdAt, updatedAt string
	err := db.QueryRow("SELECT BIN_TO_UUID(id), name, BIN_TO_UUID(user_id), file_path, file_name, BIN_TO_UUID(os_id), BIN_TO_UUID(plugin_id), DATE_FORMAT(created_at, '%Y-%m-%dT%H:%i:%sZ'), DATE_FORMAT(updated_at, '%Y-%m-%dT%H:%i:%sZ') FROM project WHERE id = UUID_TO_BIN(?)", id).Scan(
		&project.ID,
		&project.Name,
		&project.UserID,
		&project.FilePath,
		&project.FileName,
		&project.OsID,
		&project.PluginID,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}
	project.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return nil, err
	}
	project.UpdatedAt, err = time.Parse(time.RFC3339, updatedAt)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func GetProjects(ctx context.Context, db *sql.DB, userID string) ([]*Project, error) {
	projects := []*Project{}
	rows, err := db.Query("SELECT BIN_TO_UUID(id), name, BIN_TO_UUID(user_id), file_path, file_name, BIN_TO_UUID(os_id), BIN_TO_UUID(plugin_id), DATE_FORMAT(created_at, '%Y-%m-%dT%H:%i:%sZ'), DATE_FORMAT(updated_at, '%Y-%m-%dT%H:%i:%sZ') FROM project WHERE user_id = UUID_TO_BIN(?)", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var project Project
		if err := rows.Scan(&project.ID, &project.Name, &project.UserID, &project.FilePath, &project.FileName, &project.OsID, &project.PluginID, &project.CreatedAt, &project.UpdatedAt); err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}
	return projects, nil
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
