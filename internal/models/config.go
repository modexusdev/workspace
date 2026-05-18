package models

type Workspace struct {
	Name     string            `json:"name"`
	Path     string            `json:"path"`
	Commands map[string]string `json:"commands"`
}

type Config struct {
	Workspaces []Workspace `json:"workspaces"`
}
