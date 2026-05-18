package cli

import (
	"os"

	"github.com/modexusdev/workspace/internal/console"
)

func RunCli() {
	// Display CLI banner
	console.ShowBanner()
	args := os.Args
	// Handle CLI arguments
	Handle(args)
}
