package cli

import (
	"github.com/modexusdev/workspace/internal/commands"
	"github.com/modexusdev/workspace/internal/console"
)

type GlobalCommand func()
type WorkspaceCommand func(workspaceName string)

var globalCommands = map[string]GlobalCommand{
	"add":        commands.AddWorkspace,
	"ls":         commands.WorkspaceLs,
	"remove-all": commands.RemoveAllWorkspaces,
	"commands":   commands.ShowCommands,
	"version":    console.Versionshow,
}

var workspaceCommands = map[string]WorkspaceCommand{
	"start":          commands.StartWorkspace,
	"add-command":    commands.AddWorkspaceCommand,
	"remove":         commands.RemoveWorkspace,
	"ls":             commands.WorkspaceDetails,
	"set-path":       commands.SetWorkspacePath,
	"edit-command":   commands.EditWorkspaceCommand,
	"remove-command": commands.RemoveWorkspaceCommand,
	"jump":           commands.JumpToPath,
	"rename":         commands.RenameWorkspace,
}

func Handle(args []string) {
	if len(args) == 1 {
		getWorkspaces()
		return
	}

	if command, exists := globalCommands[args[1]]; exists {
		command()
		return
	}

	if len(args) == 2 {
		commands.StartWorkspace(args[1])
		return
	}

	workspaceName := args[1]
	workspaceCommand := args[2]

	if command, exists := workspaceCommands[workspaceCommand]; exists {
		command(workspaceName)
		return
	}

	commands.RunWorkspaceCommand(workspaceName, workspaceCommand)
}
