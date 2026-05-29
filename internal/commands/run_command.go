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

	workspace, found := findWorkspace(
		configData.Workspaces,
		workspaceName,
	)

	if !found {
		console.PrintError("Workspace not found.")
		return
	}

	command, found := findWorkspaceCommand(workspace.Commands, commandName)
	if !found || strings.TrimSpace(command) == "" {
		console.PrintError("Command not found.")
		return
	}

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
}
