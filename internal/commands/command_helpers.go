package commands

import (
	"bufio"
	"fmt"
	"sort"
	"strings"

	"github.com/modexusdev/workspace/internal/console"
)

func askCommandValue(reader *bufio.Reader) string {
	for {
		fmt.Printf("%sCommand:%s ", console.Cyan, console.Reset)

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "" {
			console.PrintError("Command is required.")
			continue
		}

		return command
	}
}

func askCommandName(reader *bufio.Reader, commands map[string]string) string {
	for {
		fmt.Printf("%sCommand name:%s ", console.Cyan, console.Reset)

		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			console.PrintError("Command name is required.")
			continue
		}

		if strings.Contains(name, " ") {
			console.PrintError("Command name cannot contain spaces.")
			continue
		}
		// Prevent duplicate command names
		exists := false

		for existingName := range commands {
			if strings.EqualFold(existingName, name) {
				exists = true
				break
			}
		}

		if exists {
			console.PrintError("Command name already exists.")
			continue
		}

		return name
	}
}

func askStartCommand(reader *bufio.Reader) string {

	for {

		fmt.Printf("%sStart command:%s ", console.Cyan, console.Reset)

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "" {
			console.PrintError("Start command is required.")
			continue
		}

		return command
	}
}

func askNewCommandValue(reader *bufio.Reader, commandName string) string {
	for {
		if strings.EqualFold(commandName, "start") {
			fmt.Printf("%sNew start command:%s ", console.Cyan, console.Reset)
		} else {
			fmt.Printf("%sNew command:%s ", console.Cyan, console.Reset)
		}

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "" {
			console.PrintError("Command is required.")
			continue
		}

		return command
	}
}

func askCommandToEdit(reader *bufio.Reader, commands map[string]string) string {
	for {
		fmt.Printf("%sCommand to edit:%s ", console.Cyan, console.Reset)

		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			console.PrintError("Command name is required.")
			continue
		}

		for existingName := range commands {
			if strings.EqualFold(existingName, name) {
				return existingName
			}
		}

		console.PrintError("Command not found.")
	}
}

func askNewCommandName(reader *bufio.Reader, commands map[string]string, oldCommandName string) string {
	for {
		fmt.Printf("%sNew command name:%s ", console.Cyan, console.Reset)

		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			console.PrintError("Command name is required.")
			continue
		}

		if strings.Contains(name, " ") {
			console.PrintError("Command name cannot contain spaces.")
			continue
		}

		if strings.EqualFold(name, "start") {
			console.PrintError("Only the default command can be named start.")
			continue
		}

		nameAlreadyExists := false

		for existingName := range commands {
			if strings.EqualFold(existingName, oldCommandName) {
				continue
			}

			if strings.EqualFold(existingName, name) {
				nameAlreadyExists = true
				break
			}
		}

		if nameAlreadyExists {
			console.PrintError("Command name already exists.")
			continue
		}

		return name
	}
}

func printWorkspaceCommands(commands map[string]string) {

	names := make([]string, 0, len(commands))

	for name := range commands {
		names = append(names, name)
	}

	// Sort command names for stable output
	sort.Strings(names)

	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println()

	for _, name := range names {

		fmt.Printf(
			" - %s%s%s: %s\n",
			console.Gold,
			name,
			console.Reset,
			commands[name],
		)
	}

	fmt.Println()
}

func askCommandToRemove(
	reader *bufio.Reader,
	commands map[string]string,
) string {

	for {

		fmt.Printf("%sCommand to remove:%s ", console.Cyan, console.Reset)

		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			console.PrintError("Command name is required.")
			continue
		}

		// Validate command before returning
		if !commandExists(commands, name) {
			console.PrintError("Command not found.")
			continue
		}

		return name
	}
}

func commandExists(commands map[string]string, name string) bool {

	for commandName := range commands {

		if strings.EqualFold(commandName, name) {
			return true
		}
	}

	return false
}
