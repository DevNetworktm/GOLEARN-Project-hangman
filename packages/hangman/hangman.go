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
		State:        "started",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}

	return g
}

func (game *Game) MakeAGuess(guess string) {

	if letterInWord(guess, game.UsedLetters) {
		game.State = "alreadyGuessed"
	} else if letterInWord(guess, game.Letters) {
		game.UsedLetters = append(game.UsedLetters, guess)
		game.State = "goodGuess"
		game.RevealLetter(guess)

		if hastWon(game.Letters, game.FoundLetters) {
			game.State = "won"
		}

	} else {
		game.UsedLetters = append(game.UsedLetters, guess)
		game.State = "badGuess"
		game.TurnsLeft -= 1

		if game.TurnsLeft == 0 {
			game.State = "lost"
		}
	}
}

func hastWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

func (game *Game) RevealLetter(guess string) {
	for i, l := range game.Letters {
		if l == guess {
			game.FoundLetters[i] = l
		}
	}
}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}
