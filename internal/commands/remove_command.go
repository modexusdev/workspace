package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
)

func RemoveWorkspaceCommand(workspaceName string) {

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

	if len(configData.Workspaces[index].Commands) == 0 {
		console.PrintError("No commands found.")
		return
	}

	// Display available commands before removal
	printWorkspaceCommands(configData.Workspaces[index].Commands)

	reader := bufio.NewReader(os.Stdin)

	commandName := askCommandToRemove(
		reader,
		configData.Workspaces[index].Commands,
	)

	// Prevent removing the default start command
	if strings.EqualFold(commandName, "start") {
		console.PrintError("Start command cannot be removed.")
		return
	}

	delete(configData.Workspaces[index].Commands, commandName)

	err = config.SaveConfig(configData)
	if err != nil {
		console.PrintError("Error saving config.")
		return
	}

	fmt.Println()
	console.PrintSuccess("Command removed successfully.")
	fmt.Println()
}
