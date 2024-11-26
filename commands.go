package main

import (
	"fmt"
	"os"
)

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
	limitURL := baseURL + "location-area/?limit=" + string(limit)
	offset := -limit

	mapf := func() error {
		offset += limit
		fullURL := limitURL + fmt.Sprintf("&offset=%d", offset)
		fmt.Println(fullURL)
		return nil
	}

	mapb := func() error {
		if offset <= 0 {
			fmt.Println("Error: Already at the first page of locations. Cannot go back further.")
			return nil
		}
		offset = max(0, offset-limit)
		fullURL := limitURL + fmt.Sprintf("&offset=%d", offset)
		fmt.Println(fullURL)
		return nil
	}
	return mapf, mapb
}

func MapbCommand() error {
	os.Exit(0)
	return nil
}
