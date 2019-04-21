package stemmer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

var testCases = []struct {
	word string
	stem string
}{
	{"runs", "run"},
	{"working", "work"},
	{"sleep", "sleep"},
}

func TestStem(t *testing.T) {
	require := require.New(t)
	for _, tc := range testCases {
		name := fmt.Sprintf("%s:%s", tc.word, tc.stem)
		t.Run(name, func(t *testing.T) {
			require.Equal(Stem(tc.word), tc.stem, "stem")
		})
	}
}
