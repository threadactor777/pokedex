package main

import (
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type Config struct {
	pokeapiClient pokeapi.Client
	commands      map[string]CliCommand
	next          string
	previous      string
}

type CliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
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

/*
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

	return result
}
*/

func cleanInput(text string) []string {
	words := strings.Fields(text)
	return words
}

func commandExit(userConfig *Config, paramter ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help(userConfig *Config, parameter ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandMap(userConfig *Config, parameter ...string) error {
	locationStruct, err := userConfig.pokeapiClient.GetLocation(userConfig.next)
	if err != nil {
		return err
	}

	userConfig.next = locationStruct.Next
	userConfig.previous = locationStruct.Previous

	for _, city := range locationStruct.Results {
		fmt.Println(city.Name)
	}

	return nil
}

func commandMapBack(userConfig *Config, parameter ...string) error {

	if userConfig.previous == "" {
		fmt.Println("You're already on the first page!")
		locationStruct, err := userConfig.pokeapiClient.GetLocation("https://pokeapi.co/api/v2/location-area/?offset=0&limit=20")
		if err != nil {
			return err
		}

		for _, city := range locationStruct.Results {
			fmt.Println(city.Name)
		}
		return nil
	}

	locationStruct, err := userConfig.pokeapiClient.GetLocation(userConfig.previous)
	if err != nil {
		return nil
	}

	userConfig.next = locationStruct.Next
	userConfig.previous = locationStruct.Previous

	for _, city := range locationStruct.Results {
		fmt.Println(city.Name)
	}

	return nil
}

func commandExplore(userConfig *Config, parameter ...string) error {
	return nil
}
