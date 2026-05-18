package commands

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/modexusdev/workspace/internal/config"
	"github.com/modexusdev/workspace/internal/console"
	"github.com/modexusdev/workspace/internal/models"
)

func RemoveWorkspace(workspaceName string) {

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

	updatedData, err := json.MarshalIndent(configData, "", "  ")
	if err != nil {
		console.PrintError("Error creating JSON.")
		return
	}

	err = os.WriteFile(config.ConfigPath, updatedData, 0644)
	if err != nil {
		console.PrintError("Error writing config.")
		return
	}

	console.PrintSuccess("Workspace removed successfully.")
}

func RemoveAllWorkspaces() {

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
		console.PrintError("No workspaces found.")
		return
	}

	// Ask before deleting all saved workspace entries
	if !confirmRemoveAll() {
		console.PrintError("Remove all cancelled.")
		return
	}

	configData.Workspaces = []models.Workspace{}

	updatedData, err := json.MarshalIndent(configData, "", "  ")
	if err != nil {
		console.PrintError("Error creating JSON.")
		return
	}

	err = os.WriteFile(config.ConfigPath, updatedData, 0644)
	if err != nil {
		console.PrintError("Error writing config.")
		return
	}

	console.PrintSuccess("All workspaces removed successfully.")
}

func confirmRemoveAll() bool {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf(
		"%sAre you sure you want to remove all saved workspaces?%s %s(y/N):%s ",
		console.Gold,
		console.Reset,
		console.Gray,
		console.Reset,
	)

	answer, err := reader.ReadString('\n')
	if err != nil {
		console.PrintError("Error reading confirmation.")
		return false
	}

	answer = strings.TrimSpace(answer)
	answer = strings.ToLower(answer)

	return answer == "y" || answer == "yes"
}
