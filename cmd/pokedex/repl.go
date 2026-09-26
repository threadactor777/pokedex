package main

import (
	"fmt"
	"math/rand"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type Config struct {
	pokeapiClient pokeapi.Client
	commands      map[string]CliCommand
	next          string
	previous      string
	areaBase      string
	pokemon       map[string]pokeapi.PokemonInformation
}

type CliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

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
		return fmt.Errorf("Failed to get Location!")
	}

	userConfig.next = locationStruct.Next
	userConfig.previous = locationStruct.Previous

	for _, city := range locationStruct.Results {
		fmt.Println(city.Name)
	}

	return nil
}

func commandExplore(userConfig *Config, parameter ...string) error {
	specificLocationStruct, err := userConfig.pokeapiClient.GetSpecificLocation(userConfig.areaBase + parameter[0])
	if err != nil {
		return fmt.Errorf("Failed to get Specific Location!")
	}

	fmt.Println("Exploring", parameter[0], "!")
	fmt.Println("Found:")

	for _, pokemon := range specificLocationStruct.PokemonEncounters {
		fmt.Println("-", pokemon.Pokemon.Name)
	}

	return nil
}

func commandCatch(userConfig *Config, parameter ...string) error {
	// Logic for checking location specific encounters (threaded) - not needed right now
	/*
		specificLocationStruct, err := userConfig.pokeapiClient.GetSpecificLocation(userConfig.areaBase + parameter[0])
		if err != nil {
			return fmt.Errorf("Failed to get Specific Location information to see Pokemon in area!")
		}

		// Check if pokemon exists in the location
		for _, encounter := range specificLocationStruct.PokemonEncounters {
			if encounter.Pokemon.Name == parameter[0] {
				pokemonInformationStruct, err := userConfig.pokeapiClient.GetPokemonInformation(parameter[0])
				if err != nil {
					fmt.Println("Error parsing encounters for area!")
					return err
				}
				catchProbability := rand.Intn(pokemonInformationStruct.BaseExperience)
				catchRoll := rand.Intn(pokemonInformationStruct.BaseExperience)

				fmt.Println("Throwing a Pokeball at", parameter[0], "...")

				if catchProbability == catchRoll {
					fmt.Println(parameter[0], "was caught!")
					userConfig.pokemon[parameter[0]] = pokemonInformationStruct
				}

			} else {
				fmt.Println(parameter[0], "escaped!")
			}
		}
		return nil
	*/
	pokemonInformationStruct, err := userConfig.pokeapiClient.GetPokemonInformation("https://pokeapi.co/api/v2/pokemon/" + parameter[0])
	if err != nil {
		fmt.Println("Error parsing encounters for area!")
		return err
	}

	catchProbability := rand.Intn(pokemonInformationStruct.BaseExperience) / 10
	catchRoll := rand.Intn(pokemonInformationStruct.BaseExperience) / 10

	fmt.Printf("Throwing a Pokeball at %v...\n", parameter[0])

	if catchProbability == catchRoll {
		fmt.Println(parameter[0], "was caught!")
		userConfig.pokemon[parameter[0]] = pokemonInformationStruct
	} else {
		fmt.Println(parameter[0], "escaped!")
	}
	return nil
}
