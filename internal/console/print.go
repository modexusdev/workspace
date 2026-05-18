package console

import "fmt"

func PrintError(message string) {
	fmt.Printf("%s%s✖ %s%s\n", Bold, Red, message, Reset)
}
func PrintSuccess(message string) {
	fmt.Printf("%s%s✔ %s%s\n", Bold, Green, message, Reset)
}
