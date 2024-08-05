package main

import (
	"fmt"
	"os"
)

func callbackHelp() {
	fmt.Println("Welcome to the Pokedex!\n \nUsage:")
	fmt.Println("help: Displays a help message\nexit: Exit the Pokedex\n map: display the names of 20 location areas in pokemon world ")
}

func callbackExit() {
	os.Exit(0)
}

func callbackMap() {
	fmt.Println("this is map command")
}
