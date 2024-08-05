package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	for {
		fmt.Print("pokedesk > ")

		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		text := input.Text()

		cleaned := clearInput(text)
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Invalid command")
		}

		command.callback()
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func()
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedesk",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			callback:    callbackMap,
		},
	}
}

func clearInput(str string) []string {
	lowercase := strings.ToLower(str)
	words := strings.Fields(lowercase)
	return words
}
