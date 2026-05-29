package commands

import (
	"strings"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
	"github.com/modexusdev/workspace/internal/models"
)

func RemoveWorkspace(workspaceName string) {

	configData, err := config.LoadConfig()
	if err != nil {
		console.PrintError("Error loading config.")
		return
	}
	var updatedWorkspaces []models.Workspace
	found := false

	// Remove matching workspace from config data
	for _, workspace := range configData.Workspaces {

		if strings.EqualFold(workspace.Name, workspaceName) {
			found = true
			continue
		}

		updatedWorkspaces = append(updatedWorkspaces, workspace)
	}

	if !found {
		console.PrintError("Workspace not found.")
		return
	}

	configData.Workspaces = updatedWorkspaces

	err = config.SaveConfig(configData)
	if err != nil {
		console.PrintError("Error saving config.")
		return
	}

	console.PrintSuccess("Workspace removed successfully.")
}

func RemoveAllWorkspaces() {

	configData, err := config.LoadConfig()
	if err != nil {
		console.PrintError("Error loading config.")
		return
	}

	if len(configData.Workspaces) == 0 {
		console.PrintError("No workspaces found.")
		return
	}

	// Ask before deleting all saved workspace entries
	if !confirmRemoveAll() {
		console.PrintError("Remove all cancelled.")
		return
	}

	configData.Workspaces = []models.Workspace{}

	err = config.SaveConfig(configData)
	if err != nil {
		console.PrintError("Error saving config.")
		return
	}

	console.PrintSuccess("All workspaces removed successfully.")
}
