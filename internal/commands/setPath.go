package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
	"github.com/modexusdev/workspace/internal/models"
)

func SetWorkspacePath(workspaceName string) {

	configData, err := config.LoadConfig()
	if err != nil {
		console.PrintError("Error loading config.")
		return
	}

	index := -1

	// Find target workspace
	for i, workspace := range configData.Workspaces {

		if strings.EqualFold(workspace.Name, workspaceName) {
			index = i
			break
		}
	}

	if index == -1 {
		console.PrintError("Workspace not found.")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	newPath := askNewWorkspacePath(
		reader,
		configData.Workspaces,
		workspaceName,
	)

	configData.Workspaces[index].Path = newPath

	err = config.SaveConfig(configData)
	if err != nil {
		console.PrintError("Error saving config.")
		return
	}

	fmt.Println()
	console.PrintSuccess("Workspace path updated successfully.")
	fmt.Println()
}

func askNewWorkspacePath(
	reader *bufio.Reader,
	workspaces []models.Workspace,
	currentWorkspaceName string,
) string {

	for {

		fmt.Printf(
			"%sNew workspace path%s %s(empty = current directory):%s ",
			console.Cyan,
			console.Reset,
			console.Gray,
			console.Reset,
		)

		path, _ := reader.ReadString('\n')
		path = strings.TrimSpace(path)

		// Use current directory if path is empty
		if path == "" {

			currentDir, err := os.Getwd()
			if err != nil {
				console.PrintError("Error getting current directory.")
				continue
			}

			path = currentDir
		}

		absPath, err := filepath.Abs(path)
		if err != nil {
			console.PrintError("Invalid path.")
			continue
		}

		info, err := os.Stat(absPath)
		if err != nil {

			if os.IsNotExist(err) {
				console.PrintError("Path does not exist.")
			} else {
				console.PrintError("Error checking path.")
			}

			continue
		}

		if !info.IsDir() {
			console.PrintError("Path must be a directory.")
			continue
		}

		cleanPath := filepath.Clean(absPath)

		// Prevent duplicate workspace paths
		if workspacePathExistsExcept(
			workspaces,
			cleanPath,
			currentWorkspaceName,
		) {
			console.PrintError("Workspace path already exists.")
			continue
		}

		return cleanPath
	}
}

func workspacePathExistsExcept(
	workspaces []models.Workspace,
	path string,
	currentWorkspaceName string,
) bool {

	cleanPath := filepath.Clean(path)

	for _, workspace := range workspaces {

		// Ignore current workspace while checking duplicates
		if strings.EqualFold(workspace.Name, currentWorkspaceName) {
			continue
		}

		existingPath, err := filepath.Abs(workspace.Path)
		if err != nil {
			continue
		}

		if filepath.Clean(existingPath) == cleanPath {
			return true
		}
	}

	return false
}
