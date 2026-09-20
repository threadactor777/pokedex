package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocation(page string) location {
	res, err := http.Get(page)
	if err != nil {
		fmt.Errorf("GET Request Error")
	}
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		fmt.Errorf("Response failed with a status code: %v", res.StatusCode)
	}
	if err != nil {
		fmt.Errorf("Read Response Body Error")
	}

	var locationResponse location
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		fmt.Errorf("Unmarshal Body Error")
	}
	return locationResponse
}
