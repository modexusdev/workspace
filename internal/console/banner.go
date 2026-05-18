package console

import (
	"fmt"
)

func ShowBanner() {

	fmt.Println()

	fmt.Printf(
		" %s%sWORKSPACE%s %sCLI%s %s%s%s\n",
		Bold,
		Violet,
		Reset,

		Gold,
		Reset,

		Gray,
		Version,
		Reset,
	)

	fmt.Printf("%s project manager & command runner%s\n",
		Gray, Reset,
	)

	fmt.Printf("%s────────────────────────────────────%s\n",
		DarkGold, Reset,
	)

	fmt.Println()
}
