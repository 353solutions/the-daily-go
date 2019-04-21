package nlp

import (
	//	"reflect"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func TestSentencize(t *testing.T) {
	text := `Clear is better than clever. Reflection is never clear. 
	Errors are values.`

	expected := []string{
		"Clear is better than clever.",
		"Reflection is never clear.",
		"Errors are values.",
	}

	out := Sentencize(text)
	require := require.New(t)
	require.EqualValues(expected, out)

	/*
		// MT: Ask how to compare?
		if !reflect.DeepEqual(expected, out) {
			t.Fatalf("%v != %v", expected, out)
		}
	*/
}

func TestTokenize(t *testing.T) {
	testCases := []struct {
		text     string
		expected []string
	}{
		// MT: []string{}
		{"", []string(nil)},
		{"hi", []string{"hi"}},
		{"HI", []string{"hi"}},
		{"Who's on first.", []string{"who", "s", "on", "first"}},
	}

	for _, tc := range testCases {
		// MT: start with tc.text
		name := tc.text
		if len(name) == 0 {
			name = "<empty>"
		}
		// MT: Start without t.Run
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			out := Tokenize(tc.text)
			require.Equal(tc.expected, out)
			/*
				if !reflect.DeepEqual(tc.expected, out) {
					t.Fatalf("%v != %v", tc.expected, out)
				}
			*/
		})
	}
}

func TestTokenizerQuick(t *testing.T) {
	fn := func(text string) bool {
		tokens := Tokenize(text)
		// MT: Sanity checks
		return len(wordRe.FindAllString(text, -1)) == len(tokens)
	}
	require.NoError(t, quick.Check(fn, nil))
}

// MT: Talk on data similar to real
var tokBenchText = `
Software engineering is what happens to programming when you add time and other
programmers.
    - Russ Cox
`

func BenchmarkTokenizer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toks := Tokenize(tokBenchText)
		if len(toks) != 16 {
			b.Fatal(len(toks))
		}
	}
}
