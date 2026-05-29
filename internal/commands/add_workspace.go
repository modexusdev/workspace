package commands

import (
	"bufio"
	"fmt"
	"os"

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
