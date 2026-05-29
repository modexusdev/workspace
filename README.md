# WORKSPACE CLI

```text
WORKSPACE CLI v1.0.3
project manager & command runner
────────────────────────────────────
```

![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)
[![Go Reference](https://pkg.go.dev/badge/github.com/modexusdev/workspace.svg)](https://pkg.go.dev/github.com/modexusdev/workspace)
![Platform](https://img.shields.io/badge/platform-linux-black)
![License](https://img.shields.io/badge/license-MIT-green)

A lightweight Linux-first workspace manager for daily terminal workflows.

`workspace` helps you quickly open projects, run development commands, manage editors, launch Docker workflows, and organize terminal-based development environments directly from your shell.

---

# Features

- Fast workspace management
- Open projects instantly
- Store custom project commands
- Run Docker commands quickly
- Start editors directly from terminal
- Supports shell aliases
- Supports shell functions
- Lightweight and local-first
- JSON-based configuration
- Built for terminal-focused workflows

---

# Why workspace?

Instead of constantly:

- switching between directories
- typing long paths
- reopening editors
- manually starting Docker
- manually starting development servers

You can manage everything with a single command.

---

# Installation

Option 1 — Download Binary

Download the latest binary from the [GitHub Releases](https://github.com/modexusdev/workspace/releases) page.


After downloading:

```bash
chmod +x workspace
mv workspace ~/go/bin/
```

Option 2 — Install with Go

```bash
go install github.com/modexusdev/workspace@latest
```

Make sure your Go `bin` directory is in your PATH:

```bash
echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bashrc
```

Reload your shell:

```bash
source ~/.bashrc
```

Verify installation:

```bash
workspace version
```

> [!NOTE]
> `workspace` is currently focused on Linux environments and requires `bash`.

---

# Configuration

`workspace` stores all data inside:

```bash
~/.workspace/config.json
```

Example configuration:

```json
{
  "workspaces": [
    {
      "name": "go-server",
      "path": "/home/modexusdev/projects/server",
      "commands": {
        "start": "code .",
        "docker": "docker compose up",
        "dev": "air"
      }
    }
  ]
}
```

---

# Commands


## Global Commands

### Show workspace overview

```bash
workspace
```

Displays all saved workspaces in a compact overview.


Useful for quickly seeing all registered projects.

---

### Show all workspaces

```bash
workspace ls
```

Displays a detailed list of all saved workspaces.

Shows:
- workspace names
- project paths
- configured commands

Useful when managing multiple projects.

---

### Add workspace

```bash
workspace add
```

Creates a new workspace entry interactively.

You will be asked for:
- workspace name
- project path
- default start command

Example:

```text
Workspace name: go-server
Workspace path: /home/modexus/dev/go-server
Start command: code .
```

The start command is executed when running:

```bash
workspace go-server
```

---

### Show version

```bash
workspace version
```

Displays the installed workspace CLI version.

Example:

```text
workspace v1.0.0
```

---

### Show all commands

```bash
workspace commands
```

Displays all available global and workspace commands.

Useful as a quick built-in help menu.

---

### Remove all saved workspace entries

```bash
workspace remove-all
```

Removes all saved workspaces from the config file.

Useful for resetting the entire workspace configuration.

---

# Workspace Commands

## Start workspace

```bash
workspace go-server
```

Runs the default `start` command of the workspace.

Equivalent to:

```bash
workspace go-server start
```

Example start commands:
- `code .`
- `npm run dev`
- `docker compose up`
- `go run .`

The command is automatically executed inside the saved workspace path.

---

## Run start command explicitly

```bash
workspace go-server start
```

Explicitly runs the configured `start` command.

Useful for scripting or clarity.

---

## Jump into workspace shell

```bash
workspace go-server jump
```

Opens a new shell directly inside the workspace directory.

Example:

```text
~/dev/go-server $
```

This allows working directly inside the project without manually changing directories.

Exit the shell with:

```bash
exit
```

---

## Show workspace details

```bash
workspace go-server ls
```

Displays detailed information about the workspace.

Shows:
- workspace name
- saved path
- start command
- all custom commands

Useful for inspecting workspace configuration.

---

## Rename workspace

```bash
workspace go-server rename
```

Example:

New workspace name: backend-server

## Remove workspace

```bash
workspace go-server remove
```

Removes the workspace entry from the config file.

This does NOT delete the actual project folder.

---

## Change workspace path

```bash
workspace go-server set-path
```

Updates the saved project path of the workspace.

Useful when projects are moved to another location.

Example:

```text
New path: /home/modexus/projects/go-server
```

---

## Add custom command

```bash
workspace go-server add-command
```

Adds a custom command to the workspace.

Example:

```text
Command name: docker
Command: docker compose up
```

Run later with:

```bash
workspace go-server docker
```

Useful for:
- docker
- logs
- tests
- builds
- editors
- scripts

---

## Edit custom command

```bash
workspace go-server edit-command
```

Edits an existing workspace command.

Useful when changing:
- docker commands
- dev scripts
- editor commands
- build processes

Example:

```text
Old command:
docker compose up

New command:
docker compose up --build
```

---

## Remove custom command

```bash
workspace go-server remove-command
```

Removes a custom command from the workspace.

The workspace itself remains unchanged.

---

## Run custom commands

```bash
workspace go-server docker
```

Runs the saved custom command inside the workspace directory.

Example:

```bash
docker compose up
```

Useful for quickly starting:
- development servers
- Docker environments
- build commands
- scripts
- editors
- logs

without manually navigating into the project directory.

inside the workspace directory.

Another example:

```bash
workspace go-server dev
```

---

# Shell Support

`workspace` supports:

- bash aliases
- shell functions
- custom terminal commands

Example alias:

```bash
alias dev-api='docker compose up api'
```

Example function:

```bash
mydev() {
  docker compose up
}
```

These commands can be stored directly inside a workspace.

---

# Daily Workflow Example

Open VSCode:

```bash
workspace go-server
```

Start Docker:

```bash
workspace go-server docker
```

Start development server:

```bash
workspace go-server dev
```

Open another project in Zed:

```bash
workspace go-server
```

No manual `cd`.
No long paths.
No repeated commands.

---

# Philosophy

`workspace` is designed for developers who work daily in the terminal.

The goal is to reduce repetitive actions like:

- switching between directories
- reopening editors
- manually starting Docker
- remembering project commands
- copying long file paths

Everything should be accessible through short and simple commands.

`workspace` stays:

- lightweight
- local-first
- terminal-focused
- fast
- simple

No cloud.
No telemetry.
No unnecessary complexity.

---

# License

MIT License
