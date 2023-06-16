package main

import (
	"fmt"
	"hangman/packages/hangman"
)

func main() {
	game := hangman.New(8, "Golang")
	fmt.Printf("%+v\n", game)
}
