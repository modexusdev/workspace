package commands

import (
	"bufio"
	"os"
	"strings"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
)

func RenameWorkspace(workspaceName string) {

	configData, err := config.LoadConfig()
	if err != nil {
		console.PrintError("Error loading config.")
		return
	}

	index := -1

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

	newName := askNewWorkspaceName(
		reader,
		configData.Workspaces,
		configData.Workspaces[index].Name,
	)

	configData.Workspaces[index].Name = newName

	err = config.SaveConfig(configData)
	if err != nil {
		console.PrintError("Error saving config.")
		return
	}

	console.PrintSuccess("Workspace renamed successfully.")
}
