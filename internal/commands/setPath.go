package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
)

func SetWorkspacePath(workspaceName string) {

	configData, err := config.LoadConfig()
	if err != nil {
		console.PrintError("Error loading config.")
		return
	}

	index := findWorkspaceIndex(
		configData.Workspaces,
		workspaceName,
	)

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
