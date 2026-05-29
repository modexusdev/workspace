package commands

import (
	"os"
	"os/exec"
	"strings"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
)

func StartWorkspace(workspaceName string) {

	configData, err := config.LoadConfig()
	if err != nil {
		console.PrintError("Error loading config.")
		return
	}

	for _, workspace := range configData.Workspaces {

		if !strings.EqualFold(workspace.Name, workspaceName) {
			continue
		}

		startCommand, exists := workspace.Commands["start"]

		if !exists || strings.TrimSpace(startCommand) == "" {
			console.PrintError("No start command found.")
			return
		}

		// Run start command inside workspace directory
		cmd := exec.Command("bash", "-lc", startCommand)
		cmd.Dir = workspace.Path

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		err = cmd.Run()
		if err != nil {
			console.PrintError("Error running start command: " + err.Error())
			return
		}

		return
	}

	console.PrintError("Workspace not found.")
}
