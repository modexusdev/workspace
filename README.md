# WORKSPACE CLI

```text
WORKSPACE CLI v1.0.0
project manager & command runner
────────────────────────────────────
```

![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)
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

go install github.com/modexusdev/workspace/cmd/workspace@latest

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

---

### Show all workspaces

```bash
workspace ls
```

Shows:
- workspace names
- paths
- commands

---

### Add workspace

```bash
workspace add
```

---

### Show version

```bash
workspace version
```

---

### Show all commands

```bash
workspace commands
```

---

### Remove all saved workspace entries

```bash
workspace remove-all
```

---

# Workspace Commands

## Start workspace

```bash
workspace go-server
```

Automatically runs the `start` command.

Equivalent to:

```bash
workspace go-server start
```

---

## Show workspace details

```bash
workspace go-server ls
```

Shows:
- workspace name
- path
- start command
- custom commands

---

## Remove workspace

```bash
workspace go-server remove
```

Removes the saved workspace entry from the config file.

---

## Change workspace path

```bash
workspace go-server set-path
```

---

## Add custom command

```bash
workspace go-server add-command
```

Example:

```text
Command name: docker
Command: docker compose up
```

---

## Edit custom command

```bash
workspace go-server edit-command
```

---

## Remove custom command

```bash
workspace go-server remove-command
```

---

## Run custom commands

```bash
workspace go-server docker
```

Runs:

```bash
docker compose up
```

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
