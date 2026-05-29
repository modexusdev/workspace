package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
)

func EditWorkspaceCommand(workspaceName string) {
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
	// Display available commands before editing
	if len(configData.Workspaces[index].Commands) == 0 {
		console.PrintError("No commands found.")
		return
	}

	printWorkspaceCommands(configData.Workspaces[index].Commands)

	reader := bufio.NewReader(os.Stdin)

	oldCommandName := askCommandToEdit(reader, configData.Workspaces[index].Commands)

	newCommandName := oldCommandName
	// Prevent renaming the default start command
	if !strings.EqualFold(oldCommandName, "start") {
		newCommandName = askNewCommandName(reader, configData.Workspaces[index].Commands, oldCommandName)
	}

	newCommandValue := askNewCommandValue(reader, oldCommandName)
	// Resolve aliases or shell commands before saving
	resolvedCommand, err := ResolveCommandOnSave(newCommandValue)
	if err != nil {
		console.PrintError(err.Error())
		return
	}
	// Remove old command entry if renamed
	if !strings.EqualFold(oldCommandName, newCommandName) {
		delete(configData.Workspaces[index].Commands, oldCommandName)
	}

	configData.Workspaces[index].Commands[newCommandName] = resolvedCommand

	err = config.SaveConfig(configData)
	if err != nil {
		console.PrintError("Error saving config.")
		return
	}

	fmt.Println()
	console.PrintSuccess("Command updated successfully.")
	fmt.Println()
}
