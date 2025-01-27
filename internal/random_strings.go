package internal

import (
	"math/rand"
	"strings"
)

var randomWords = []string{
	"humpty",
	"dumpty",
	"Horsey",
	"donkey",
	"yikes",
	"monkey",
	"Coo",
	"scooby",
	"dooby",
	"vanilla",
	"237",
	"Monkey",
}

func randomWord() string {
	return randomWords[rand.Intn(len(randomWords))]
}

func randomString(n int, joiner string) string {
	b := make([]string, n)

	for i := range b {
		b[i] = randomWord()
	}

	return strings.Join(b, joiner)
}

func randomAnything() string {
	size := rand.Intn(2) + 1
	return randomWord() + "/" + randomString(size, "-")
}

func randomUrlPath() string {
	return randomAnything()
}

func randomUserAgent() string {
	return randomAnything()
}

func randomFileName() string {
	return randomString(4, "_")
}

func randomFileNameWithPrefix(prefix string) string {
	return prefix + randomFileName()
}

func randomFileContent() string {
	return randomString(100, " ")
}
