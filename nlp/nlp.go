package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
)

var (
	// Punctuation followed by whitespace
	// Two or more newlines
	sentenceEndRe = regexp.MustCompile("([\\.!?][[:space:]])|(\n\n+)")
	wordRe        = regexp.MustCompile("[[:alpha:]]+")
)

func sentencize(text []byte) [][]byte {
	var sentences [][]byte
	start := 0
	for start < len(text) {
		loc := sentenceEndRe.FindIndex(text[start:])
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

func tokenize(text []byte) [][]byte {
	words := wordRe.FindAll(text, -1)
	tokens := make([][]byte, len(words)) // MT: 0 and profile?
	for i, w := range words {
		tokens[i] = bytes.ToLower(w)
	}
	return tokens
}

type sentence struct {
	text   []byte
	tokens [][]byte
	score  int
}

type byScore []*sentence

func (s byScore) Len() int               { return len(s) }
func (s byScore) Less(i int, j int) bool { return s[i].score < s[j].score }
func (s byScore) Swap(i int, j int)      { s[i], s[j] = s[j], s[i] }

func score(tokens [][]byte, freqs map[string]int) int {
	score := 0
	for _, tok := range tokens {
		score += freqs[string(tok)]
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func summarize(text []byte, count int) [][]byte {
	freqs := make(map[string]int)
	var sents []*sentence // MT: Make in length?

	for _, s := range sentencize(text) {
		tokens := tokenize(s)
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
	var summary [][]byte
	for _, sent := range sents[:count] {
		summary = append(summary, sent.text)
	}

	return summary

}

func main() {
	// https://en.wikinews.org/wiki/Turkey_recalls_ambassador_to_Sweden_over_%22genocide%22_vote
	text, err := ioutil.ReadFile("testdata/1.txt")
	if err != nil {
		panic(err)
	}

	summary := summarize(text, 3)
	for _, s := range summary {
		fmt.Println(string(s))
		fmt.Println("==========================")
	}
}
