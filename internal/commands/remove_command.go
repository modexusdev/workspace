package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
	"github.com/modexusdev/workspace/internal/models"
)

func RemoveWorkspaceCommand(workspaceName string) {

	data, err := os.ReadFile(config.ConfigPath)
	if err != nil {
		console.PrintError("Error reading config.")
		return
	}

	var configData models.Config

	err = json.Unmarshal(data, &configData)
	if err != nil {
		console.PrintError("Error parsing config.")
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

	updatedData, err := json.MarshalIndent(configData, "", "  ")
	if err != nil {
		console.PrintError("Error creating JSON.")
		return
	}

	err = os.WriteFile(config.ConfigPath, updatedData, 0644)
	if err != nil {
		console.PrintError("Error writing config.")
		return
	}

	fmt.Println()
	console.PrintSuccess("Command removed successfully.")
	fmt.Println()
}

func printWorkspaceCommands(commands map[string]string) {

	names := make([]string, 0, len(commands))

	for name := range commands {
		names = append(names, name)
	}

	// Sort command names for stable output
	sort.Strings(names)

	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println()

	for _, name := range names {

		fmt.Printf(
			" - %s%s%s: %s\n",
			console.Gold,
			name,
			console.Reset,
			commands[name],
		)
	}

	fmt.Println()
}

func askCommandToRemove(
	reader *bufio.Reader,
	commands map[string]string,
) string {

	for {

		fmt.Printf("%sCommand to remove:%s ", console.Cyan, console.Reset)

		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			console.PrintError("Command name is required.")
			continue
		}

		// Validate command before returning
		if !commandExists(commands, name) {
			console.PrintError("Command not found.")
			continue
		}

		return name
	}
}

func commandExists(commands map[string]string, name string) bool {

	for commandName := range commands {

		if strings.EqualFold(commandName, name) {
			return true
		}
	}

	return false
}
