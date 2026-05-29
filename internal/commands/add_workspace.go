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

func AddWorkspace() {

	reader := bufio.NewReader(os.Stdin)

	configData, err := config.LoadConfig()
	if err != nil {
		console.PrintError("Error loading config.")
		return
	}

	name := askWorkspaceName(reader, configData.Workspaces)
	path := askWorkspacePath(reader, configData.Workspaces)
	startCommand := askStartCommand(reader)

	// Resolve aliases or shell commands before saving
	resolvedCommand, err := ResolveCommandOnSave(startCommand)
	if err != nil {
		console.PrintError(err.Error())
		return
	}

	newWorkspace := models.Workspace{
		Name: name,
		Path: path,
		Commands: map[string]string{
			"start": resolvedCommand,
		},
	}

	configData.Workspaces = append(configData.Workspaces, newWorkspace)

	err = config.SaveConfig(configData)
	if err != nil {
		console.PrintError("Error saving config.")
		return
	}

	fmt.Println()
	console.PrintSuccess("Workspace added successfully.")
	fmt.Println()
}

func askWorkspaceName(reader *bufio.Reader, workspaces []models.Workspace) string {

	for {

		fmt.Printf("%sWorkspace name:%s ", console.Cyan, console.Reset)

		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			console.PrintError("Workspace name is required.")
			continue
		}
		// Prevent duplicate workspace names
		if workspaceNameExists(workspaces, name) {
			console.PrintError("Workspace name already exists.")
			continue
		}

		return name
	}
}

func askWorkspacePath(reader *bufio.Reader, workspaces []models.Workspace) string {

	for {

		fmt.Printf("%sWorkspace path%s %s(empty = current directory):%s ",
			console.Cyan, console.Reset,
			console.Gray, console.Reset,
		)

		path, _ := reader.ReadString('\n')
		path = strings.TrimSpace(path)

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
		if workspacePathExists(workspaces, cleanPath) {
			console.PrintError("Workspace path already exists.")
			continue
		}

		return cleanPath
	}
}

func askStartCommand(reader *bufio.Reader) string {

	for {

		fmt.Printf("%sStart command:%s ", console.Cyan, console.Reset)

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "" {
			console.PrintError("Start command is required.")
			continue
		}

		return command
	}
}

func workspaceNameExists(workspaces []models.Workspace, name string) bool {

	for _, workspace := range workspaces {

		if strings.EqualFold(workspace.Name, name) {
			return true
		}
	}

	return false
}

func workspacePathExists(workspaces []models.Workspace, path string) bool {

	cleanPath := filepath.Clean(path)

	for _, workspace := range workspaces {

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
