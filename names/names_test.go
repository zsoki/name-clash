package names

import (
	"fmt"
	"testing"
)

func TestLevenshtein(t *testing.T) {
	distance := levenshteinDistance([]byte("Anna"), []byte("Hanna"))
	fmt.Printf("Levenshtein distance: %v\n", distance)
}
