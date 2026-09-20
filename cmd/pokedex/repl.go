package main

import (
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type Config struct {
	commands map[string]CliCommand
	next     string
	previous string
}

type CliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

/*
var commands = map[string]CliCommand{
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
*/

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

func commandExit(userConfig *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help(userConfig *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandMap(userConfig *Config) error {
	locationStruct := pokeapi.GetLocation(userConfig.next)

	userConfig.next = locationStruct.Next
	userConfig.previous = locationStruct.Previous

	for _, city := range locationStruct.Results {
		fmt.Println(city.Name)
	}

	return nil
}

func commandMapBack(userConfig *Config) error {

	if userConfig.previous == "" {
		fmt.Printf("You're already on the first page!")
		return nil
	}

	locationStruct := pokeapi.GetLocation(userConfig.previous)

	userConfig.next = locationStruct.Next
	userConfig.previous = locationStruct.Previous

	for _, city := range locationStruct.Results {
		fmt.Println(city.Name)
	}

	return nil
}
