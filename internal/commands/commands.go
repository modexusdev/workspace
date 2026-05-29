package commands

import (
	"fmt"

	"github.com/modexusdev/workspace/internal/console"
)

func ShowCommands() {
	fmt.Println()
	fmt.Printf("%sGlobal Commands%s\n\n", console.Violet, console.Reset)

	fmt.Printf(" %sworkspace%s                           show workspace overview\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace ls%s                        list all workspaces\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace add%s                       add new workspace\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace commands%s                  show all available commands\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace version%s                   show CLI version\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace remove-all%s                remove all saved workspace entries\n", console.Gold, console.Reset)

	fmt.Println()
	fmt.Printf("%sWorkspace Commands%s\n\n", console.Violet, console.Reset)

	fmt.Printf(" %sworkspace <name>%s                    run default start command\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace <name> start%s              run start command explicitly\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace <name> jump%s               open shell in workspace path\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace <name> ls%s                 show workspace details\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace <name> rename%s             rename saved workspace entry\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace <name> add-command%s        add custom command\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace <name> edit-command%s       edit existing command\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace <name> remove-command%s     remove custom command\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace <name> set-path%s           update workspace path\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace <name> <command>%s          run custom command\n", console.Gold, console.Reset)
	fmt.Printf(" %sworkspace <name> remove%s             remove saved workspace entry\n", console.Gold, console.Reset)

	fmt.Println()
}
