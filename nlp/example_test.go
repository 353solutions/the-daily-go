package nlp_test

import (
	"fmt"

	"nlp"
)

func ExampleTokenize() {
	s := "Hi, how are you feeling today?"
	tokens := nlp.Tokenize(s)
	fmt.Println(tokens)
	// Output: [hi feel today]
}

func ExampleSentencize() {
	text := `
	Hi there! How are you feeling today?
	I'm good, how about you?
	`

	for _, sent := range nlp.Sentencize(text) {
		fmt.Println(sent)
	}
	// Output:
	// Hi there!
	// How are you feeling today?
	// I'm good, how about you?
}
