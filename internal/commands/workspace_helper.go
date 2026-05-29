package commands

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/modexusdev/workspace/internal/console"
	"github.com/modexusdev/workspace/internal/models"
)

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

func workspaceNameExists(workspaces []models.Workspace, name string) bool {

	for _, workspace := range workspaces {

		if strings.EqualFold(workspace.Name, name) {
			return true
		}
	}

	return false
}
func confirmRemoveAll() bool {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf(
		"%sAre you sure you want to remove all saved workspaces?%s %s(y/N):%s ",
		console.Gold,
		console.Reset,
		console.Gray,
		console.Reset,
	)

	answer, err := reader.ReadString('\n')
	if err != nil {
		console.PrintError("Error reading confirmation.")
		return false
	}

	answer = strings.TrimSpace(answer)
	answer = strings.ToLower(answer)

	return answer == "y" || answer == "yes"
}

func confirm(message string) bool {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf(
		"%s%s%s %s(y/N):%s ",
		console.Gold,
		message,
		console.Reset,
		console.Gray,
		console.Reset,
	)

	answer, err := reader.ReadString('\n')
	if err != nil {
		console.PrintError("Error reading confirmation.")
		return false
	}

	answer = strings.TrimSpace(answer)
	answer = strings.ToLower(answer)

	return answer == "y" || answer == "yes"
}

func findWorkspaceCommand(commands map[string]string, commandName string) (string, bool) {

	// Find command case-insensitive
	for existingName, command := range commands {

		if strings.EqualFold(existingName, commandName) {
			return command, true
		}
	}

	return "", false
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

func findWorkspaceIndex(workspaces []models.Workspace, workspaceName string) int {
	for i, workspace := range workspaces {
		if strings.EqualFold(workspace.Name, workspaceName) {
			return i
		}
	}

	return -1
}
func findWorkspace(
	workspaces []models.Workspace,
	workspaceName string,
) (*models.Workspace, bool) {
	for i := range workspaces {
		if strings.EqualFold(workspaces[i].Name, workspaceName) {
			return &workspaces[i], true
		}
	}

	return nil, false
}
func askNewWorkspaceName(
	reader *bufio.Reader,
	workspaces []models.Workspace,
	oldWorkspaceName string,
) string {

	for {
		console.PrintSuccess("New workspace name: ")

		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			console.PrintError("Workspace name is required.")
			continue
		}

		if strings.Contains(name, " ") {
			console.PrintError("Workspace name cannot contain spaces.")
			continue
		}

		nameAlreadyExists := false

		for _, workspace := range workspaces {
			if strings.EqualFold(workspace.Name, oldWorkspaceName) {
				continue
			}

			if strings.EqualFold(workspace.Name, name) {
				nameAlreadyExists = true
				break
			}
		}

		if nameAlreadyExists {
			console.PrintError("Workspace name already exists.")
			continue
		}

		return name
	}
}
