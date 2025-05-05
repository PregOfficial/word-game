package game

import (
	"strings"
	"time"
)

type Game struct {
	Id          string
	SecretWord  string
	MaxAttempts int
	Attempts    int
	IsOver      bool
	Won         bool
	Guesses     [][]LetterResult
	CreatedAt   time.Time
	ExpiresAt   time.Time
}

type Color string

const (
	ColorGray   Color = "gray"
	ColorYellow Color = "yellow"
	ColorGreen  Color = "green"
)

type LetterResult struct {
	Letter string `json:"letter"`
	Color  Color  `json:"color"`
}

func NewGame(id string, words []string) Game {
	createdAt := time.Now()

	return Game{
		Id:          id,
		SecretWord:  SelectRandomWord(words),
		MaxAttempts: 5,
		Attempts:    0,
		CreatedAt:   createdAt,
		ExpiresAt:   createdAt.Add(5 * time.Minute),
	}
}

func (g *Game) Guess(guess string) {
	result := g.checkGuess(guess)
	g.Attempts++

	if g.isCorrectGuess(guess) {
		g.Won = true
		g.IsOver = true
	}
	if g.Attempts >= g.MaxAttempts {
		g.IsOver = true
	}

	g.Guesses = append(g.Guesses, result)
}

func (g Game) checkGuess(guess string) []LetterResult {
	secretWordLetters := strings.Split(g.SecretWord, "")
	letters := strings.Split(guess, "")
	result := make([]LetterResult, 5)
	used := make([]bool, 5)

	for i, letter := range letters {
		if strings.Compare(letter, secretWordLetters[i]) == 0 {
			result[i] = LetterResult{
				Letter: letter,
				Color:  ColorGreen,
			}
			used[i] = true
		}
	}

	for i := range letters {
		if result[i].Color != "" {
			continue
		}
		for j := range secretWordLetters {
			if !used[j] && letters[i] == secretWordLetters[j] {
				result[i] = LetterResult{
					Letter: letters[i],
					Color:  ColorYellow,
				}
				used[j] = true
				break
			}
		}
		if result[i].Color == "" {
			result[i] = LetterResult{
				Letter: letters[i],
				Color:  ColorGray,
			}
		}
	}

	return result
}

func (g Game) isCorrectGuess(guess string) bool {
	return strings.Compare(g.SecretWord, guess) == 0
}
