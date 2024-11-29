package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type PokeLocationBody struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var baseURL = "https://pokeapi.co/api/v2/"

func ExitCommand() error {
	os.Exit(0)
	return nil
}

func HelpCommand() error {
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MapCommands() (func() error, func() error) {
	const limit = 20
	limitURL := baseURL + "location-area/?limit=" + fmt.Sprintf("%d", limit)
	offset := -limit

	getPokeLocations := func(url string) ([]string, error) {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		if res.StatusCode > 299 {
			return nil,
				fmt.Errorf("response failed with status code: %d and\nbody: %s",
					res.StatusCode, body)
		}

		var pokeLocations []string
		var PokeLocObj PokeLocationBody
		err = json.Unmarshal(body, &PokeLocObj)
		if err != nil {
			return nil, err
		}
		for _, location := range PokeLocObj.Results {
			pokeLocations = append(pokeLocations, location.Name)
		}

		return pokeLocations, nil
	}

	printStringList := func(list []string) {
		for _, item := range list {
			fmt.Println(item)
		}
	}

	mapf := func() error {
		offset += limit
		fullURL := limitURL + fmt.Sprintf("&offset=%d", offset)
		pokeLocations, err := getPokeLocations(fullURL)
		if err != nil {
			return err
		}
		printStringList(pokeLocations)
		return nil
	}

	mapb := func() error {
		if offset <= 0 {
			fmt.Println("Error: Already at the first page of locations. Cannot go back further.")
			return nil
		}
		offset = max(0, offset-limit)
		fullURL := limitURL + fmt.Sprintf("&offset=%d", offset)
		pokeLocations, err := getPokeLocations(fullURL)
		if err != nil {
			return err
		}
		printStringList(pokeLocations)
		return nil
	}
	return mapf, mapb
}

func MapbCommand() error {
	os.Exit(0)
	return nil
}
