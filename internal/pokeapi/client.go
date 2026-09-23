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
