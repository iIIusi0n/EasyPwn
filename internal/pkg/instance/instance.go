package instance

import (
	"context"
	"database/sql"
	"easypwn/internal/pkg/project"
	"easypwn/internal/pkg/util"
	"fmt"
	"time"
)

type Instance struct {
	ID          string
	ProjectID   string
	ContainerID string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewInstance(ctx context.Context, db *sql.DB, projectID string) (*Instance, error) {
	project, err := project.GetProject(ctx, db, projectID)
	if err != nil {
		return nil, err
	}

	imageName := fmt.Sprintf("easypwn/%s:%s", project.OsID, project.PluginID)

	containerName := util.CreateInstanceName()
	containerID, err := createContainer(ctx, cli, containerName, imageName, project.FilePath, true)
	if err != nil {
		return nil, err
	}

	err = startContainer(ctx, cli, containerID)
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var instanceID string
	result := tx.QueryRow("INSERT INTO instance (project_id, container_id) VALUES (?, ?) RETURNING id", projectID, containerID)
	err = result.Scan(&instanceID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return &Instance{
		ID:          instanceID,
		ProjectID:   projectID,
		ContainerID: containerID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func GetInstance(ctx context.Context, db *sql.DB, id string) (*Instance, error) {
	instance := &Instance{}
	err := db.QueryRow("SELECT * FROM instance WHERE id = $1", id).Scan(
		&instance.ID,
		&instance.ProjectID,
		&instance.ContainerID,
		&instance.CreatedAt,
		&instance.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func (i *Instance) Stop() error {
	return stopContainer(context.Background(), cli, i.ContainerID)
}

func (i *Instance) Delete(ctx context.Context, db *sql.DB) error {
	err := i.Stop()
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM instance WHERE id = $1", i.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (i *Instance) GetLogs(ctx context.Context, db *sql.DB, limit int) (string, error) {
	var logs string
	err := db.QueryRow("SELECT log FROM instance_log WHERE instance_id = $1 ORDER BY created_at DESC LIMIT $2", i.ID, limit).Scan(&logs)
	if err != nil {
		return "", err
	}
	return logs, nil
}