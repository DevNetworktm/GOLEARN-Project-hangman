package main

import (
	"fmt"
	"hangman/packages/hangman"
	"os"
	"strings"
)

func main() {
	game := hangman.New(8, "Golang")
	fmt.Printf("%+v\n", game)

	hangman.DrawWelcome()

	guess := ""

	for {
		hangman.Draw(game, guess)

		switch game.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Error reading input: %s\n", err)
			os.Exit(1)
		}

		guess = strings.ToUpper(l)
	}

}
