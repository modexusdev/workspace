package main

import (
	"github.com/modexusdev/workspace/internal/cli"
	"github.com/modexusdev/workspace/internal/config"
)

func main() {

	// Init Workspace config and required directories
	config.InitializeConfig()

	// Start CLI handler
	cli.RunCli()
}
