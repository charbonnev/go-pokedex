package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    ExitCommand,
		},
		"help": {
			name:        "help",
			description: "Get a list of commands",
			callback: func() error {
				for _, command := range getCommands() {
					fmt.Printf("%s: %s\n", command.name, command.description)
				}
				return nil
			},
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world. Call again for the next 20 locations",
			callback: func() error {
				os.Exit(0)
				return nil
			},
		},
		"mapb": {
			name:        "mapb",
			description: "pairs with map, displays the names of the 20 previous location areas in the Pokemon world",
			callback: func() error {
				os.Exit(0)
				return nil
			},
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
