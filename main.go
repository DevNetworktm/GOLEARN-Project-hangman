package main

import (
	"fmt"
	"hangman/packages/dictionary"
	"hangman/packages/hangman"
	"os"
	"strings"
)

func main() {
	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load dictionary: %v\n", err)
		os.Exit(1)
	}

	game, err := hangman.New(8, dictionary.GetRandomWord())

	if err != nil {
		fmt.Printf("Error starting a new game: %v\n", err)
		os.Exit(1)
	}

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
		game.MakeAGuess(guess)
	}

}
