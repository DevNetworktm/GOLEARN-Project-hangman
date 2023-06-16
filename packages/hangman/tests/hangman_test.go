package hangman__test

import (
	"hangman/packages/hangman"
	"testing"
)

func TestInvalideWord(t *testing.T) {
	_, err := hangman.New(3, "")
	if err == nil {
		t.Errorf("Error should be returned when using an empty word")
	}
}

func TestGameGoodGuess(t *testing.T) {
	g, _ := hangman.New(3, "bobe")
	g.MakeAGuess("B")
	validState(t, "goodGuess", g.State)
}

func TestGameAlreadyGuessed(t *testing.T) {
	g, _ := hangman.New(3, "bob2")
	g.MakeAGuess("B")
	g.MakeAGuess("B")
	validState(t, "alreadyGuessed", g.State)
}

func TestGameBadGuess(t *testing.T) {
	g, _ := hangman.New(3, "bobe")
	g.MakeAGuess("A")
	validState(t, "badGuess", g.State)
}

func TestGameWon(t *testing.T) {
	g, _ := hangman.New(3, "bobe")
	g.MakeAGuess("B")
	g.MakeAGuess("O")
	g.MakeAGuess("E")
	validState(t, "won", g.State)
}

func TestGameLost(t *testing.T) {
	g, _ := hangman.New(3, "bobe")
	g.MakeAGuess("D")
	g.MakeAGuess("I")
	g.MakeAGuess("M")
	validState(t, "lost", g.State)
}

func validState(t *testing.T, expectedState, actualState string) bool {
	if expectedState != actualState {
		t.Errorf("State should be '%s'. got=%s", expectedState, actualState)
		return false
	}
	return true
}
