package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

func cleanInput(input string) []string {
	result := strings.Fields(input)
	for i, word := range result {
		result[i] = strings.ToLower(word)
	}
	return result
}

func main() {
	cliMap := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display this help",
			callback:    commandHelp,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		cmd, ok := cliMap[cleanInput(scanner.Text())[0]]
		if ok {
			cmd.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}
