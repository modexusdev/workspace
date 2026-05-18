package cli

import (
	"github.com/modexusdev/workspace/internal/commands"
	"github.com/modexusdev/workspace/internal/console"
)

func Handle(args []string) {

	if len(args) == 1 {
		getWorkspaces()
		return
	}

	// Handle global commands
	switch args[1] {

	case "add":
		commands.AddWorkspace()
		return

	case "ls":
		commands.WorkspaceLs()
		return
	case "remove-all":
		commands.RemoveAllWorkspaces()
		return
	case "version":
		return
	case "commands":
		commands.ShowCommands()
		return
	}
	// Run default workspace command
	if len(args) == 2 {
		commands.StartWorkspace(args[1])
		return
	}

	if len(args) < 3 {
		console.PrintError("Missing workspace command.")
		return
	}

	workspaceName := args[1]
	workspaceCommand := args[2]

	// Handle workspace commands
	switch workspaceCommand {

	case "start":
		commands.StartWorkspace(workspaceName)
	case "add-command":
		commands.AddWorkspaceCommand(workspaceName)
	case "remove":
		commands.RemoveWorkspace(workspaceName)
	case "ls":
		commands.WorkspaceDetails(workspaceName)
	case "set-path":
		commands.SetWorkspacePath(workspaceName)
	case "edit-command":
		commands.EditWorkspaceCommand(workspaceName)
	case "remove-command":
		commands.RemoveWorkspaceCommand(workspaceName)

	default:
		commands.RunWorkspaceCommand(workspaceName, workspaceCommand)
	}
}
