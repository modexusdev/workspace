package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
	"github.com/modexusdev/workspace/internal/models"
)

func getWorkspaces() {

	data, err := os.ReadFile(config.ConfigPath)
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	var config models.Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		console.PrintError(fmt.Sprintf("Error reading config: %v", err))
		return
	}

	// Render workspace overview
	showWorkspaceList(config.Workspaces)
}

func showWorkspaceList(workspaces []models.Workspace) {

	fmt.Printf("%sWorkspaces%s %s(%d)%s\n\n",
		console.Violet, console.Reset,
		console.Gold, len(workspaces), console.Reset,
	)

	if len(workspaces) == 0 {
		fmt.Printf(" %sNo workspaces found.%s\n\n", console.Gray, console.Reset)

		fmt.Printf(
			" %sworkspace add%s    create your first workspace\n\n",
			console.Gold,
			console.Reset,
		)
	} else {
		for i, workspace := range workspaces {
			fmt.Printf(
				" %s%02d%s %s▸%s %s%s%s\n",
				console.Gold, i+1, console.Reset,
				console.Gray, console.Reset,
				console.Violet, workspace.Name, console.Reset,
			)
		}

		fmt.Println()
	}
	// Show quick help entry
	fmt.Printf(
		"%sHelp%s\n\n",
		console.Gray,
		console.Reset,
	)

	fmt.Printf(
		" %sworkspace commands%s    show all available commands\n",
		console.Gold,
		console.Reset,
	)

	fmt.Println()
}
