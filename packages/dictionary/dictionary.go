package dictionary

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

var words = make([]string, 0, 50)

func Load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		wordsByLine := strings.Split(line, ",")
		for _, word := range wordsByLine {
			if word == "" {
				continue
			}
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func GetRandomWord() string {
	rand.Seed(time.Now().Unix())
	return words[rand.Intn(len(words))]
}
