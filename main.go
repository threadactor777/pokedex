package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
				commandExit()
			case "help":
				help()
			default:
				fmt.Printf("Unknown command")
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Errorf("Error:", err)
		}

	}
}
