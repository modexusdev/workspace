package config

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	HomeDir      string
	WorkspaceDir string
	ConfigPath   string
)

func InitializeConfig() {

	// Resolve user home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting user home directory: %v\n", err)
		return
	}

	HomeDir = homeDir
	// Setup workspace Paths
	WorkspaceDir = filepath.Join(HomeDir, ".workspace")
	ConfigPath = filepath.Join(WorkspaceDir, "config.json")

	ensureWorkspaceDir()
	createConfigFile()
}

func ensureWorkspaceDir() {

	// Create workspace directory if it dose not exist
	err := os.MkdirAll(WorkspaceDir, 0755)
	if err != nil {
		fmt.Printf("Error creating workspace directory: %v\n", err)
		return
	}
}

func createConfigFile() {
	// Check if config file already exists
	_, err := os.Stat(ConfigPath)

	if err == nil {
		return
	}

	if !os.IsNotExist(err) {
		fmt.Printf("Error checking config file: %v\n", err)
		return
	}

	defaultConfig := `{
  "workspaces": []
}`
	// Create defualt config file
	err = os.WriteFile(ConfigPath, []byte(defaultConfig), 0644)
	if err != nil {
		fmt.Printf("Error writing config file: %v\n", err)
		return
	}
}
