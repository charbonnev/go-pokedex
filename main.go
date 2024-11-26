package main

import (
	"fmt"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	mapForward, mapBack := MapCommands()
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    ExitCommand,
		},
		"help": {
			name:        "help",
			description: "Get a list of commands",
			callback:    HelpCommand,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas in the Pokemon world. Call again for the next 20.",
			callback:    mapForward,
		},
		"mapb": {
			name:        "mapb",
			description: "Pairs with map, displays the 20 previous location areas.",
			callback:    mapBack,
		},
	}
}

func main() {
	commands := getCommands()
	fmt.Printf("Pokedex > ")
	for {
		var input string
		fmt.Scanln(&input)
		cleanInput := strings.ToLower(input)
		if command, ok := commands[cleanInput]; ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command, type \"help\" for a list of commands")
		}
		fmt.Printf("Pokedex > ")
	}
}
