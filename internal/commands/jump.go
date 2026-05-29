package commands

import (
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
)

func JumpToPath(workspaceName string) {
	configData, err := config.LoadConfig()
	if err != nil {
		console.PrintError("Error loading config.")
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
