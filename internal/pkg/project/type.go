package project

import "easypwn/internal/data"

type Os struct {
	ID   string
	Name string
}

type Plugin struct {
	ID   string
	Name string
}

func GetOsNameFromID(id string) (string, error) {
	db := data.GetDB()
	var name string
	err := db.QueryRow("SELECT name FROM project_os WHERE id = UUID_TO_BIN(?)", id).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func GetOsIDFromName(name string) (string, error) {
	db := data.GetDB()
	var id string
	err := db.QueryRow("SELECT BIN_TO_UUID(id) FROM project_os WHERE name = ?", name).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func GetPluginNameFromID(id string) (string, error) {
	db := data.GetDB()
	var name string
	err := db.QueryRow("SELECT name FROM project_plugin WHERE id = UUID_TO_BIN(?)", id).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func GetPluginIDFromName(name string) (string, error) {
	db := data.GetDB()
	var id string
	err := db.QueryRow("SELECT BIN_TO_UUID(id) FROM project_plugin WHERE name = ?", name).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func GetOsList() ([]Os, error) {
	db := data.GetDB()
	rows, err := db.Query("SELECT BIN_TO_UUID(id), name FROM project_os")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	oss := []Os{}
	for rows.Next() {
		var os Os
		if err := rows.Scan(&os.ID, &os.Name); err != nil {
			return nil, err
		}
		oss = append(oss, os)
	}
	return oss, nil
}

func GetPluginList() ([]Plugin, error) {
	db := data.GetDB()
	rows, err := db.Query("SELECT BIN_TO_UUID(id), name FROM project_plugin")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	plugins := []Plugin{}
	for rows.Next() {
		var plugin Plugin
		if err := rows.Scan(&plugin.ID, &plugin.Name); err != nil {
			return nil, err
		}
		plugins = append(plugins, plugin)
	}
	return plugins, nil
}
