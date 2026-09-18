package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
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
}

func cleanInput(text string) []string {
	var result []string
	var temp string
	for _, character := range text {
		if string(character) == " " {
			if len(temp) > 1 {
				result = append(result, strings.ToLower(temp))
				temp = ""
			}
			continue
		}
		temp = strings.ToLower(temp + string(character))
	}

	if len(temp) > 1 {
		result = append(result, strings.ToLower(temp))
	}

	return result
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}
