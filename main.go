package main

import (
	"bufio"
	"fmt"
	"os"
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
			input := cleanInput(scanner.Text())
			// Print first word
			fmt.Println("Your command was:", input[0])
		}

		if err := scanner.Err(); err != nil {
			fmt.Errorf("Error:", err)
		}

	}
}
