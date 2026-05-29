package commands

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ResolveCommandOnSave(input string) (string, error) {

	input = strings.TrimSpace(input)

	if input == "" {
		return "", errors.New("command cannot be empty")
	}

	parts := strings.Fields(input)

	first := parts[0]
	rest := parts[1:]

	// Check if command exists in PATH
	if _, err := exec.LookPath(first); err == nil {
		return input, nil
	}

	aliases := LoadShellAliases()

	// Resolve shell alias
	if aliasValue, ok := aliases[first]; ok {

		resolvedParts := append([]string{aliasValue}, rest...)

		return strings.Join(resolvedParts, " "), nil
	}

	functions := LoadShellFunctions()

	// Resolve shell function
	if functionBody, ok := functions[first]; ok {

		callParts := append([]string{first}, rest...)
		callCommand := strings.Join(callParts, " ")

		return functionBody + "\n\n" + callCommand, nil
	}

	return "", errors.New("command not found: " + first)
}

func LoadShellAliases() map[string]string {

	aliases := make(map[string]string)

	home, err := os.UserHomeDir()
	if err != nil {
		return aliases
	}

	files := []string{
		filepath.Join(home, ".bash_aliases"),
		filepath.Join(home, ".bashrc"),
		filepath.Join(home, ".zshrc"),
	}

	// Load aliases from supported shell config files
	for _, file := range files {
		readAliasesFromFile(file, aliases)
	}

	return aliases
}

func LoadShellFunctions() map[string]string {

	functions := make(map[string]string)

	home, err := os.UserHomeDir()
	if err != nil {
		return functions
	}

	files := []string{
		filepath.Join(home, ".bash_aliases"),
		filepath.Join(home, ".bashrc"),
		filepath.Join(home, ".zshrc"),
	}

	// Load shell functions from supported shell config files
	for _, file := range files {
		readFunctionsFromFile(file, functions)
	}

	return functions
}

func readAliasesFromFile(path string, aliases map[string]string) {

	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		if !strings.HasPrefix(line, "alias ") {
			continue
		}

		name, value, ok := parseAliasLine(line)

		if ok {
			aliases[name] = value
		}
	}
}

func readFunctionsFromFile(path string, functions map[string]string) {

	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	for i := 0; i < len(lines); i++ {

		line := strings.TrimSpace(lines[i])

		name, ok := parseFunctionStart(line)
		if !ok {
			continue
		}

		var block []string

		braceCount := 0
		started := false

		// Collect complete function block
		for j := i; j < len(lines); j++ {

			currentLine := lines[j]

			block = append(block, currentLine)

			braceCount += strings.Count(currentLine, "{")
			braceCount -= strings.Count(currentLine, "}")

			if strings.Contains(currentLine, "{") {
				started = true
			}

			if started && braceCount == 0 {

				functions[name] = strings.Join(block, "\n")

				i = j

				break
			}
		}
	}
}

func parseAliasLine(line string) (string, string, bool) {

	line = strings.TrimSpace(line)

	if !strings.HasPrefix(line, "alias ") {
		return "", "", false
	}

	line = strings.TrimPrefix(line, "alias ")
	line = strings.TrimSpace(line)

	parts := strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		return "", "", false
	}

	name := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])

	if name == "" || value == "" {
		return "", "", false
	}

	if len(value) >= 2 {
		first := value[0]
		last := value[len(value)-1]

		if (first == '\'' && last == '\'') || (first == '"' && last == '"') {
			value = value[1 : len(value)-1]
		}
	}

	value = strings.TrimSpace(value)

	if value == "" {
		return "", "", false
	}

	return name, value, true
}

func parseFunctionStart(line string) (string, bool) {

	line = strings.TrimSpace(line)

	// Match: function myfunc() {
	if strings.HasPrefix(line, "function ") {

		line = strings.TrimPrefix(line, "function ")
		line = strings.TrimSpace(line)

		parts := strings.Fields(line)

		if len(parts) == 0 {
			return "", false
		}

		name := strings.TrimSuffix(parts[0], "()")
		name = strings.TrimSpace(name)

		if name != "" && strings.Contains(line, "{") {
			return name, true
		}
	}

	// Match: myfunc() {
	if strings.Contains(line, "()") && strings.Contains(line, "{") {

		name := strings.Split(line, "()")[0]
		name = strings.TrimSpace(name)

		if name != "" {
			return name, true
		}
	}

	return "", false
}
