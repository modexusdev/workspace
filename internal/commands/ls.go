package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
	"github.com/modexusdev/workspace/internal/models"
)

func WorkspaceLs() {
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

	if len(configData.Workspaces) == 0 {
		fmt.Printf(
			"%sNo workspaces found.%s\n",
			console.Gray,
			console.Reset,
		)
		return
	}

	fmt.Printf("\nWorkspaces (%d)\n\n", len(configData.Workspaces))

	for i, workspace := range configData.Workspaces {
		fmt.Printf(
			"%s%02d%s ▸ %s%s%s\n",
			console.DarkBlue,
			i+1,
			console.Reset,

			console.Violet,
			workspace.Name,
			console.Reset,
		)
		fmt.Printf(
			"   %sPath:%s %s%s%s\n",
			console.Gold,
			console.Reset,

			console.Gray,
			workspace.Path,
			console.Reset,
		)

		if len(workspace.Commands) == 0 {
			fmt.Println("   Commands: none")
		} else {
			fmt.Println("   Commands:")
			// Sort command names for stable output
			commandNames := make([]string, 0, len(workspace.Commands))
			for name := range workspace.Commands {
				commandNames = append(commandNames, name)
			}

			sort.Strings(commandNames)

			for _, name := range commandNames {
				fmt.Printf(
					"   - %s%s%s: %s\n",
					console.Gold,
					name,
					console.Reset,
					workspace.Commands[name],
				)
			}
		}

		fmt.Println()
	}
}

func WorkspaceDetails(workspaceName string) {

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
	// Find workspace by name
	for _, workspace := range configData.Workspaces {

		if !strings.EqualFold(workspace.Name, workspaceName) {
			continue
		}

		fmt.Println()

		fmt.Printf(
			"%sWorkspace:%s %s%s%s\n\n",
			console.Gold,
			console.Reset,

			console.Violet,
			workspace.Name,
			console.Reset,
		)

		fmt.Printf(
			"%sPath:%s %s%s%s\n\n",
			console.Gold,
			console.Reset,

			console.Gray,
			workspace.Path,
			console.Reset,
		)

		if len(workspace.Commands) == 0 {

			fmt.Printf(
				"%sNo commands found.%s\n\n",
				console.Gray,
				console.Reset,
			)

			return
		}

		fmt.Printf(
			"%sCommands:%s\n\n",
			console.Gold,
			console.Reset,
		)

		commandNames := make([]string, 0, len(workspace.Commands))
		// Sort command names for stable output
		for name := range workspace.Commands {
			commandNames = append(commandNames, name)
		}

		sort.Strings(commandNames)

		for _, name := range commandNames {

			fmt.Printf(
				" - %s%s%s: %s\n",
				console.Gold,
				name,
				console.Reset,

				workspace.Commands[name],
			)
		}

		fmt.Println()

		return
	}

	console.PrintError("Workspace not found.")
}
