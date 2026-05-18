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

func StartWorkspace(workspaceName string) {

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
