package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
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
