package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func ExitCommand(cfg *Config) error {
	os.Exit(0)
	return nil
}

func HelpCommand(cfg *Config) error {
	for _, command := range getCommands(cfg) {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func MapCommands(cfg *Config) (mapf, mapb func(*Config) error) {
	const limit = 20
	initUrl := baseURL + "location-area/?limit=" + fmt.Sprintf("%d", limit)
	if cfg.Next == nil {
		cfg.Next = &initUrl
	}

	getPokeLocations := func(url string) (pokeLocations []string, nextURL *string, err error) {
		res, err := http.Get(url)
		if err != nil {
			return nil, nil, err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, nil, err
		}
		if res.StatusCode > 299 {
			return nil, nil,
				fmt.Errorf("response failed with status code: %d and\nbody: %s",
					res.StatusCode, body)
		}

		var PokeLocObj PokeLocationBody
		err = json.Unmarshal(body, &PokeLocObj)
		if err != nil {
			return nil, nil, err
		}
		for _, location := range PokeLocObj.Results {
			pokeLocations = append(pokeLocations, location.Name)
		}

		return pokeLocations, PokeLocObj.Next, err
	}

	printStringList := func(list []string) {
		for _, item := range list {
			fmt.Println(item)
		}
	}

	// mapf = map forward - print the next 20 location areas
	mapf = func(cfg *Config) error {
		pokeLocations, nextURL, err := getPokeLocations(*cfg.Next)
		if err != nil {
			return err
		}
		cfg.Previous = cfg.Next
		cfg.Next = nextURL
		printStringList(pokeLocations)
		return nil
	}

	// mapb = map backward - print the previous 20 location areas
	mapb = func(cfg *Config) error {
		if cfg.Previous == nil {
			return fmt.Errorf("error: Already at the first page of locations. Cannot go back further")
		}
		pokeLocations, nextURL, err := getPokeLocations(*cfg.Previous)
		if err != nil {
			return err
		}
		cfg.Previous = cfg.Next
		cfg.Next = nextURL
		printStringList(pokeLocations)
		return nil
	}
	return mapf, mapb
}
