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

type PokemonInformation struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	PastAbilities []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Abilities []struct {
			IsHidden bool        `json:"is_hidden"`
			Slot     int         `json:"slot"`
			Ability  interface{} `json:"ability"`
		} `json:"abilities"`
	} `json:"past_abilities"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	HeldItems              []interface{} `json:"held_items"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt int `json:"level_learned_at"`
			VersionGroup   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			Order interface{} `json:"order"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		Other struct {
			Home struct {
				FrontShiny       string      `json:"front_shiny"`
				FrontFemale      interface{} `json:"front_female"`
				FrontDefault     string      `json:"front_default"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
			} `json:"home"`
			Showdown struct {
				BackShiny        string      `json:"back_shiny"`
				BackFemale       interface{} `json:"back_female"`
				FrontShiny       string      `json:"front_shiny"`
				BackDefault      string      `json:"back_default"`
				FrontFemale      interface{} `json:"front_female"`
				FrontDefault     string      `json:"front_default"`
				BackShinyFemale  interface{} `json:"back_shiny_female"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
			} `json:"showdown"`
			DreamWorld struct {
				FrontFemale  interface{} `json:"front_female"`
				FrontDefault string      `json:"front_default"`
			} `json:"dream_world"`
			OfficialArtwork struct {
				FrontShiny   string `json:"front_shiny"`
				FrontDefault string `json:"front_default"`
			} `json:"official-artwork"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				Yellow struct {
					BackGray         string `json:"back_gray"`
					FrontGray        string `json:"front_gray"`
					BackDefault      string `json:"back_default"`
					FrontDefault     string `json:"front_default"`
					BackTransparent  string `json:"back_transparent"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"yellow"`
				RedBlue struct {
					BackGray         string `json:"back_gray"`
					FrontGray        string `json:"front_gray"`
					BackDefault      string `json:"back_default"`
					FrontDefault     string `json:"front_default"`
					BackTransparent  string `json:"back_transparent"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"red-blue"`
			} `json:"generation-i"`
			GenerationV struct {
				Icons struct {
					Animated struct {
						FrontDefault string `json:"front_default"`
					} `json:"animated"`
					FrontDefault string `json:"front_default"`
				} `json:"icons"`
				BlackWhite struct {
					Animated struct {
						BackShiny        string      `json:"back_shiny"`
						BackFemale       interface{} `json:"back_female"`
						FrontShiny       string      `json:"front_shiny"`
						BackDefault      string      `json:"back_default"`
						FrontFemale      interface{} `json:"front_female"`
						FrontDefault     string      `json:"front_default"`
						BackShinyFemale  interface{} `json:"back_shiny_female"`
						FrontShinyFemale interface{} `json:"front_shiny_female"`
					} `json:"animated"`
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationIi struct {
				Gold struct {
					BackShiny        string `json:"back_shiny"`
					FrontShiny       string `json:"front_shiny"`
					BackDefault      string `json:"back_default"`
					FrontDefault     string `json:"front_default"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackShiny        string `json:"back_shiny"`
					FrontShiny       string `json:"front_shiny"`
					BackDefault      string `json:"back_default"`
					FrontDefault     string `json:"front_default"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"silver"`
				Crystal struct {
					Animated struct {
						FrontShiny   string `json:"front_shiny"`
						FrontDefault string `json:"front_default"`
					} `json:"animated"`
					BackShiny             string `json:"back_shiny"`
					FrontShiny            string `json:"front_shiny"`
					BackDefault           string `json:"back_default"`
					FrontDefault          string `json:"front_default"`
					BackTransparent       string `json:"back_transparent"`
					FrontTransparent      string `json:"front_transparent"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
				} `json:"crystal"`
			} `json:"generation-ii"`
			GenerationIv struct {
				Platinum struct {
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"platinum"`
				DiamondPearl struct {
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackShiny        string      `json:"back_shiny"`
					BackFemale       interface{} `json:"back_female"`
					FrontShiny       string      `json:"front_shiny"`
					BackDefault      string      `json:"back_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
			} `json:"generation-iv"`
			GenerationIx struct {
				ScarletViolet struct {
					FrontFemale  interface{} `json:"front_female"`
					FrontDefault string      `json:"front_default"`
				} `json:"scarlet-violet"`
			} `json:"generation-ix"`
			GenerationVi struct {
				XY struct {
					FrontShiny       string      `json:"front_shiny"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"x-y"`
				OmegarubyAlphasapphire struct {
					FrontShiny       string      `json:"front_shiny"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     string      `json:"front_default"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
			} `json:"generation-vi"`
			GenerationIii struct {
				Emerald struct {
					FrontShiny   string `json:"front_shiny"`
					FrontDefault string `json:"front_default"`
				} `json:"emerald"`
				RubySapphire struct {
					BackShiny    string `json:"back_shiny"`
					FrontShiny   string `json:"front_shiny"`
					BackDefault  string `json:"back_default"`
					FrontDefault string `json:"front_default"`
				} `json:"ruby-sapphire"`
				FireredLeafgreen struct {
					BackShiny    string `json:"back_shiny"`
					FrontShiny   string `json:"front_shiny"`
					BackDefault  string `json:"back_default"`
					FrontDefault string `json:"front_default"`
				} `json:"firered-leafgreen"`
			} `json:"generation-iii"`
			GenerationVii struct {
				Icons struct {
					FrontFemale  interface{} `json:"front_female"`
					FrontDefault string      `json:"front_default"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontShiny       interface{} `json:"front_shiny"`
					FrontFemale      interface{} `json:"front_female"`
					FrontDefault     interface{} `json:"front_default"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontFemale  interface{} `json:"front_female"`
					FrontDefault string      `json:"front_default"`
				} `json:"icons"`
				BrilliantDiamondShiningPearl struct {
					FrontFemale  interface{} `json:"front_female"`
					FrontDefault string      `json:"front_default"`
				} `json:"brilliant-diamond-shining-pearl"`
			} `json:"generation-viii"`
		} `json:"versions"`
		BackShiny        string      `json:"back_shiny"`
		BackFemale       interface{} `json:"back_female"`
		FrontShiny       string      `json:"front_shiny"`
		BackDefault      string      `json:"back_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontDefault     string      `json:"front_default"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	PastStats []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		Stats []struct {
			BaseStat int `json:"base_stat"`
			Effort   int `json:"effort"`
			Stat     struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"stat"`
		} `json:"stats"`
	} `json:"past_stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	PastTypes []interface{} `json:"past_types"`
}

func (c *Client) GetPokemonInformation(page string) (PokemonInformation, error) {
	if val, exists := c.cache.Get(page); exists {
		fmt.Println("(using cached data)")
		var pokemonInformationResponse PokemonInformation
		err := json.Unmarshal(val, &pokemonInformationResponse)
		if err != nil {
			return PokemonInformation{}, err
		}
		return pokemonInformationResponse, nil
	}

	res, err := http.Get(page)
	if err != nil {
		return PokemonInformation{}, fmt.Errorf("GET Request Error")
	}
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return PokemonInformation{}, fmt.Errorf("Response failed with a status code: %v", res.StatusCode)
	}
	if err != nil {
		return PokemonInformation{}, fmt.Errorf("Read Response Body Error")
	}

	var pokemonInformationResponse PokemonInformation
	err = json.Unmarshal(body, &pokemonInformationResponse)
	if err != nil {
		return PokemonInformation{}, fmt.Errorf("Unmarshal Body Error")
	}

	c.cache.Add(page, body)

	return pokemonInformationResponse, nil
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
