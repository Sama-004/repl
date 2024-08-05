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
		command := cleaned[0]
		switch command {
		case "exit":
			os.Exit(0)
		case "help":
			fmt.Println("Welcome to the Pokedex!\n \nUsage:")
			fmt.Println("help: Displays a help message\nexit: Exit the Pokedex")
		default:
			fmt.Println("Invalid command")
		}
	}
}

func clearInput(str string) []string {
	lowercase := strings.ToLower(str)
	words := strings.Fields(lowercase)
	return words
}
