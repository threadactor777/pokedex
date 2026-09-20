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
			case "map":
				commandMap(currentConfig)
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
		"map": {
			name:        "map",
			description: "Displays 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in the Pokemon world",
			callback:    commandMapBack,
		},
	},
	next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	previous: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
}

func main() {
	// Start REPL loop
	repl(&userConfig)
}
