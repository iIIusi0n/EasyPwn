package project

import "easypwn/internal/data"

func GetOsNameFromID(id string) (string, error) {
	db := data.GetDB()
	var name string
	err := db.QueryRow("SELECT name FROM project_os WHERE id = UUID_TO_BIN($1)", id).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func GetOsIDFromName(name string) (string, error) {
	db := data.GetDB()
	var id string
	err := db.QueryRow("SELECT BIN_TO_UUID(id) FROM project_os WHERE name = $1", name).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func GetPluginNameFromID(id string) (string, error) {
	db := data.GetDB()
	var name string
	err := db.QueryRow("SELECT name FROM project_plugin WHERE id = UUID_TO_BIN($1)", id).Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}

func GetPluginIDFromName(name string) (string, error) {
	db := data.GetDB()
	var id string
	err := db.QueryRow("SELECT BIN_TO_UUID(id) FROM project_plugin WHERE name = $1", name).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
