// Package nlp providwes some NLP functions
package nlp

//go:generate sh -c "go run gen_stop.go < stop_words.txt > stop_words.go"
//go:generate gofmt -w stop_words.go

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	// MT
	//	"nlp/stemmer"
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
	add := func(s string) {
		s = strings.TrimSpace(s)
		if len(s) > 0 {
			sentences = append(sentences, s)
		}
	}

	for start < len(text) {
		loc := sentenceEndRe.FindStringIndex(text[start:])
		if loc == nil {
			if start < len(text) {
				add(text[start:])
			}
			break
		}
		add(text[start : start+loc[1]])
		start += loc[1]
	}

	return sentences
}

// Tokenize will split text to list of tokens
func Tokenize(text string) []string {
	words := wordRe.FindAllString(text, -1)
	var tokens []string
	// MT: Optimize (~5%)
	// tokens := make([]string, 0, 20)
	for _, w := range words {
		token := strings.ToLower(w)
		/* MT
		token = stemmer.Stem(token)
		if StopWords[token] {
			continue
		}
		*/
		tokens = append(tokens, token)
	}
	return tokens
}

type sentence struct {
	text   string
	tokens []string
	score  int
}

type byScore []*sentence

func (s byScore) Len() int               { return len(s) }
func (s byScore) Less(i int, j int) bool { return s[i].score < s[j].score }
func (s byScore) Swap(i int, j int)      { s[i], s[j] = s[j], s[i] }

func score(tokens []string, freqs map[string]int) int {
	score := 0
	for _, tok := range tokens {
		score += freqs[tok]
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Summarize will extract "count" most significat stentences from text
func Summarize(text string, count int) []string {
	freqs := make(map[string]int)
	var sents []*sentence // MT: Make in length?

	for _, s := range Sentencize(text) {
		tokens := Tokenize(s)
		for _, token := range tokens {
			freqs[string(token)]++
		}
		sents = append(sents, &sentence{s, tokens, 0})
	}

	for _, s := range sents {
		s.score = score(s.tokens, freqs)
		fmt.Println(s.score)
	}

	sort.Sort(sort.Reverse(byScore(sents)))
	count = min(count, len(sents))
	var summary []string
	for _, sent := range sents[:count] {
		summary = append(summary, sent.text)
	}

	return summary
}
