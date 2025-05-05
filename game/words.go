package game

import (
	"bufio"
	"math/rand"
	"os"
)

func LoadWords(filename string) ([]string, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func SelectRandomWord(words []string) string {
	if len(words) < 1 {
		return ""
	}
	return words[rand.Intn(len(words))]
}
