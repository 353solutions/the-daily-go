package stemmer

import (
	"strings"
)

var (
	suffixes = []string{"s", "ing"}
)

// Stem returns the stemmed version of word
func Stem(word string) string {
	for _, suffix := range suffixes {
		if strings.HasSuffix(word, suffix) {
			return word[:len(word)-len(suffix)]
		}
	}
	return word
}
