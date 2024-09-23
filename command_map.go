package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type NameUrl struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type RespShallowLocations struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []NameUrl `json:"results"`
}

type OrchastratePagination struct {
	currentUrl string
	nextUrl    string
	data       string
}

func commandLocationf() error {
	url := "https://pokeapi.co/api/v2/location/"
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Printf("errror in get res")
	}

	location := RespShallowLocations{}
	if err := json.NewDecoder(res.Body).Decode(&location); err != nil {
		return err
	}

	locationFormatted := locationFormat(location)
	fmt.Println(locationFormatted)

	return nil
}

func locationFormat(locations RespShallowLocations) string {
	locationString := []string{}
	for _, val := range locations.Results {
		locationString = append(locationString, val.Name)
	}
	locationFmt := strings.Join(locationString, "\n")
	return locationFmt
}
