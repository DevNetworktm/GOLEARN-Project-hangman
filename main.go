package main

import (
	"fmt"
	"hangman/packages/hangman"
	"os"
)

func main() {
	game := hangman.New(8, "Golang")
	fmt.Printf("%+v\n", game)

	l, err := hangman.ReadGuess()
	if err != nil {
		fmt.Printf("Error reading input: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("You guessed: %s\n", l)
}
