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

func askCommandToEdit(reader *bufio.Reader, commands map[string]string) string {
	for {
		fmt.Printf("%sCommand to edit:%s ", console.Cyan, console.Reset)

		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			console.PrintError("Command name is required.")
			continue
		}

		for existingName := range commands {
			if strings.EqualFold(existingName, name) {
				return existingName
			}
		}

		console.PrintError("Command not found.")
	}
}

func askNewCommandName(reader *bufio.Reader, commands map[string]string, oldCommandName string) string {
	for {
		fmt.Printf("%sNew command name:%s ", console.Cyan, console.Reset)

		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			console.PrintError("Command name is required.")
			continue
		}

		if strings.Contains(name, " ") {
			console.PrintError("Command name cannot contain spaces.")
			continue
		}

		if strings.EqualFold(name, "start") {
			console.PrintError("Only the default command can be named start.")
			continue
		}

		nameAlreadyExists := false

		for existingName := range commands {
			if strings.EqualFold(existingName, oldCommandName) {
				continue
			}

			if strings.EqualFold(existingName, name) {
				nameAlreadyExists = true
				break
			}
		}

		if nameAlreadyExists {
			console.PrintError("Command name already exists.")
			continue
		}

		return name
	}
}

func askNewCommandValue(reader *bufio.Reader, commandName string) string {
	for {
		if strings.EqualFold(commandName, "start") {
			fmt.Printf("%sNew start command:%s ", console.Cyan, console.Reset)
		} else {
			fmt.Printf("%sNew command:%s ", console.Cyan, console.Reset)
		}

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "" {
			console.PrintError("Command is required.")
			continue
		}

		return command
	}
}
