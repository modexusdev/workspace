package commands

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
	"github.com/modexusdev/workspace/internal/models"
)

func RunWorkspaceCommand(workspaceName string, commandName string) {

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

	for _, workspace := range configData.Workspaces {

		if !strings.EqualFold(workspace.Name, workspaceName) {
			continue
		}

		command, found := findWorkspaceCommand(workspace.Commands, commandName)

		if !found || strings.TrimSpace(command) == "" {
			console.PrintError("Command not found.")
			return
		}

		// Run command inside workspace directory
		cmd := exec.Command("bash", "-lc", command)
		cmd.Dir = workspace.Path

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		err = cmd.Run()
		if err != nil {
			console.PrintError("Error running command: " + err.Error())
			return
		}

		return
	}

	console.PrintError("Workspace not found.")
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
