package commands

import (
	"os"
	"os/exec"
	"strings"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
)

func RunWorkspaceCommand(workspaceName string, commandName string) {

	configData, err := config.LoadConfig()
	if err != nil {
		console.PrintError("Error loading config.")
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
