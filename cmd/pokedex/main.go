package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"time"
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
			input := cleanInput(scanner.Text())

			/* Need to create function to handle blank entry
			-
			-
			-
			-
			*/

			// Check if they want to exit
			if len(input) == 0 {
				fmt.Println("You didnt enter anything!")
				continue
			}

			if len(input) > 1 {
				switch input[0] {
				case "catch":
					commandCatch(currentConfig, input[1])
				case "explore":
					commandExplore(currentConfig, input[1])
				}
				continue
			}
			switch input[0] {
			case "exit":
				commandExit(currentConfig)
			case "help":
				help(currentConfig)
			case "map":
				commandMap(currentConfig)
			case "mapb":
				commandMapBack(currentConfig)
			case "catch":
				fmt.Println("No Pokemon Specified!")
			case "explore":
				fmt.Println("No Location Specified!")
			default:
				fmt.Println("Unknown command")
			}

		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error: %v", err)
		}
	}
}

var userConfig = Config{
	pokeapiClient: pokeapi.NewClient(5 * time.Minute),
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
		"explore": {
			name:        "explore",
			description: "Displays the pokemon in a specific area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to cath an encountered pokemon",
			callback:    commandCatch,
		},
	},
	next:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	previous: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
	areaBase: "https://pokeapi.co/api/v2/location-area/",
	pokemon:  map[string]pokeapi.PokemonInformation{},
	//need to initialize above?????
}

func main() {
	// Start REPL loop
	repl(&userConfig)
}
