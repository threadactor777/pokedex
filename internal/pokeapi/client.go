package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

type Client struct {
	cache *pokecache.Cache
}

func NewClient(interval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(interval),
	}
}

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type SpecificLocation struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				Chance   int `json:"chance"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				ConditionValues []interface{} `json:"condition_values"`
				PokemonDetails  interface{}   `json:"pokemon_details"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetSpecificLocation(page string) (SpecificLocation, error) {
	if val, exists := c.cache.Get(page); exists {
		fmt.Println("(using cached data)")
		var specificLocationResponse SpecificLocation
		err := json.Unmarshal(val, &specificLocationResponse)
		if err != nil {
			return SpecificLocation{}, err
		}
		return specificLocationResponse, nil
	}

	res, err := http.Get(page)
	if err != nil {
		return SpecificLocation{}, fmt.Errorf("GET Request Error")
	}
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return SpecificLocation{}, fmt.Errorf("Response failed with a status code: %v", res.StatusCode)
	}
	if err != nil {
		return SpecificLocation{}, fmt.Errorf("Read Response Body Error")
	}

	var specificLocationResponse SpecificLocation
	err = json.Unmarshal(body, &specificLocationResponse)
	if err != nil {
		return SpecificLocation{}, fmt.Errorf("Unmarshal Body Error")
	}

	c.cache.Add(page, body)

	return specificLocationResponse, nil
}

func (c *Client) GetLocation(page string) (Location, error) {
	if val, exists := c.cache.Get(page); exists {
		fmt.Println("(using cached data)")
		var locationResponse Location
		err := json.Unmarshal(val, &locationResponse)
		if err != nil {
			return Location{}, err
		}
		return locationResponse, nil
	}

	res, err := http.Get(page)
	if err != nil {
		return Location{}, fmt.Errorf("GET Request Error")
	}
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return Location{}, fmt.Errorf("Response failed with a status code: %v", res.StatusCode)
	}
	if err != nil {
		return Location{}, fmt.Errorf("Read Response Body Error")
	}

	var locationResponse Location
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return Location{}, fmt.Errorf("Unmarshal Body Error")
	}

	c.cache.Add(page, body)

	return locationResponse, nil
}
