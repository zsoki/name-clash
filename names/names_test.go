package names

import (
	"testing"
)

func TestLevenshtein(t *testing.T) {
	compare(t, "Anna", "Hanna", 1)
	compare(t, "Dorottya", "Dorotea", 2)
}

func compare(t *testing.T, name1, name2 string, expected int) {
	actual := levenshteinDistance(convertToSymbolic(name1), convertToSymbolic(name2))
	if expected != actual {
		t.Error("Levenshtein failed", expected, "!=", actual)
	}
}
