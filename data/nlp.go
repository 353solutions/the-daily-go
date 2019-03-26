package nlp

import (
	"regexp"
	"strings"
)

var (
	// Punctuation followed by whitespace
	// Two or more newlines
	sentenceEndRe = regexp.MustCompile("([\\.!?][[:space:]])|(\n\n+)")
	wordRe        = regexp.MustCompile("[[:alpha:]]+")
)

// Sentencize will split text to list of sentences
func Sentencize(text string) []string {
	var sentences []string
	start := 0
	for start < len(text) {
		loc := sentenceEndRe.FindStringIndex(text[start:])
		if loc == nil {
			if start < len(text) {
				sentences = append(sentences, text[start:])
			}
			break
		}
		sentences = append(sentences, text[start:start+loc[1]])
		start += loc[1]
	}

	return sentences
}

// Tokenize will split text to list of tokens
func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	var tokens []string
	for _, w := range words {
		token := strings.ToLower(w)
		tokens = append(tokens, token)
	}
	return tokens
}
