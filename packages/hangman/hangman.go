package hangman

import "strings"

type Game struct {
	State        string   // State of the game
	Letters      []string // Letters guessed
	FoundLetters []string // Letters found
	UsedLetters  []string // Letters used
	TurnsLeft    int      // Turns left
}

func New(turns int, word string) *Game {
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))

	for i, _ := range letters {
		found[i] = "_"
	}

	g := &Game{
		State:        "Started",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}

	return g
}
