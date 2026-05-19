package commands

import (
	"encoding/json"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
	"github.com/modexusdev/workspace/internal/models"
)

func JumpToPath(workspaceName string) {
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

		err = os.Chdir(workspace.Path)
		if err != nil {
			console.PrintError("Could not jump to workspace path.")
			return
		}

		shell, err := exec.LookPath("bash")
		if err != nil {
			console.PrintError("bash not found.")
			return
		}

		err = syscall.Exec(shell, []string{shell}, os.Environ())
		if err != nil {
			console.PrintError("Could not open shell.")
			return
		}
	}

	console.PrintError("Workspace not found.")
}
