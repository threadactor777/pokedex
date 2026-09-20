package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl(currentConfig *Config) {
	// Create scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Infinite for loop
	for {
		// Print cursor thingy
		fmt.Print("Pokedex > ")

		// Scan tokens
		if scanner.Scan() {
			input := strings.Join(cleanInput(scanner.Text()), "")

			// Check if they want to exit
			switch input {
			case "exit":
				commandExit(currentConfig)
			case "help":
				help(currentConfig)
			default:
				fmt.Printf("Unknown command")
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Errorf("Error:", err)
		}

	}
}

var userConfig = Config{
	commands: map[string]CliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    help,
		},
	},
}

func main() {
	repl(&userConfig)
}
