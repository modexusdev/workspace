package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
)

func AddWorkspaceCommand(workspaceName string) {
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

	commandName := askCommandName(reader, configData.Workspaces[index].Commands)
	commandValue := askCommandValue(reader)

	// Resolve aliases or shell commands before saving
	resolvedCommand, err := ResolveCommandOnSave(commandValue)
	if err != nil {
		console.PrintError(err.Error())
		return
	}
	// Initialize commands map if missing
	if configData.Workspaces[index].Commands == nil {
		configData.Workspaces[index].Commands = make(map[string]string)
	}

	configData.Workspaces[index].Commands[commandName] = resolvedCommand

	err = config.SaveConfig(configData)
	if err != nil {
		console.PrintError("Error saving config.")
		return
	}

	fmt.Println()
	console.PrintSuccess("Command added successfully.")
	fmt.Println()
}

func askCommandName(reader *bufio.Reader, commands map[string]string) string {
	for {
		fmt.Printf("%sCommand name:%s ", console.Cyan, console.Reset)

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
		// Prevent duplicate command names
		exists := false

		for existingName := range commands {
			if strings.EqualFold(existingName, name) {
				exists = true
				break
			}
		}

		if exists {
			console.PrintError("Command name already exists.")
			continue
		}

		return name
	}
}

func askCommandValue(reader *bufio.Reader) string {
	for {
		fmt.Printf("%sCommand:%s ", console.Cyan, console.Reset)

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "" {
			console.PrintError("Command is required.")
			continue
		}

		return command
	}
}
