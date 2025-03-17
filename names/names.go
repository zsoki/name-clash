package names

import (
	"fmt"
	"slices"
	"strings"
)

const matchLength int = 5

type Name struct {
	Text     string
	Symbolic []byte
	Edges    []*Edge
}

type Edge struct {
	Names    [2]*Name
	Distance int
}

func (e *Edge) neighborOf(from *Name) *Name {
	for _, to := range e.Names {
		if from != to {
			return to
		}
	}
	return nil
}

func Run(filename, example string, maxDistance int) ([]*Name, []*Edge) {
	lines := make(chan string)
	go ReadLines(filename, lines)

	edges := make([]*Edge, 0)
	names := make([]*Name, 0)

	for line := range lines {
		newName := Name{Text: line, Symbolic: convertToSymbolic(line)}

		// Calculate Levenshtein distance with exiting names, connect both of them with an edge each
		for _, existingName := range names {
			distance := levenshteinDistance(newName.Symbolic, existingName.Symbolic)
			if distance <= maxDistance {
				edge := Edge{[2]*Name{existingName, &newName}, distance}
				edges = append(edges, &edge)
				newName.Edges = append(newName.Edges, &edge)
				existingName.Edges = append(existingName.Edges, &edge)
			}
		}

		names = append(names, &newName)
	}

	{ // DEBUG LOGGING
		fmt.Printf("All names: %v\n", len(names))
		fmt.Printf("%v example:\n", example)

		exampleMatches := func(name *Name) bool { return name.Text == example }
		existingIdx := slices.IndexFunc(names, exampleMatches)

		if existingIdx != -1 {
			name := names[existingIdx]
			for _, edge := range name.Edges {
				fmt.Printf("%v (distance: %v)\n", edge.neighborOf(name).Text, edge.Distance)
			}
		}
	} // DEBUG LOGGING

	var round = 1
	var players = len(names) - (len(names) % 2)
	var allMatches = 0
	var matches int

	for players > 1 {
		matches = players / 2
		allMatches += matches
		fmt.Printf("Round %v: %v players, %v matches\n", round, players, matches)
		players = players / 2
		players = players - players%2
		round++
	}

	fmt.Printf("All matches: %v\n", allMatches)
	fmt.Printf("If one match takes %v seconds, the tournament lasts for %v\n", matchLength, allMatches*matchLength)

	return names, edges
}

func levenshteinDistance(source, target []byte) int {
	tableRows := len(source) + 1
	tableCols := len(target) + 1

	table := make([][]int, tableRows)
	for row := range table {
		table[row] = make([]int, tableCols)
	}

	for row := 1; row < tableRows; row++ {
		const firstCol = 0
		table[row][firstCol] = row
	}

	for col := 1; col < tableCols; col++ {
		const firstRow = 0
		table[firstRow][col] = col
	}

	for col := 1; col < tableCols; col++ {
		for row := 1; row < tableRows; row++ {

			substitutionCost := 1
			sourceChar := source[row-1]
			targetChar := target[col-1]
			if sourceChar == targetChar {
				substitutionCost = 0
			}

			// deletion, insertion, substitution
			edits := [3]int{
				table[row-1][col] + 1,
				table[row][col-1] + 1,
				table[row-1][col-1] + substitutionCost,
			}

			table[row][col] = slices.Min(edits[:])
		}
	}

	// printTable(table)
	return table[tableRows-1][tableCols-1]
}

var multiGraphs = []string{
	"cs", "dz", "dzs", "gy", "ly", "ny", "sz", "ty", "zs",
	"bb", "cc", "ccs", "dd", "ddz", "ddzs", "ff", "gg", "ggy", "hh", "jj", "kk", "ll", "lly", "mm", "nn", "nny", "pp", "rr", "ss", "ssz", "tt", "tty", "vv", "zz", "zzs",
}

var diacritics = []rune{
	'á', 'é', 'í', 'ó', 'ö', 'ő', 'ú', 'ü', 'ű',
}

func convertToSymbolic(input string) []byte {
	input = strings.ToLower(input)
	for idx, multiGraph := range multiGraphs {
		input = strings.ReplaceAll(input, multiGraph, string(rune(idx)))
	}
	for idx, diacritic := range diacritics {
		input = strings.ReplaceAll(input, string(diacritic), string(rune(len(multiGraphs)+idx)))
	}
	return []byte(input)
}

func printTable(table [][]int) {
	for row := 0; row < len(table); row++ {
		for col := 0; col < len(table[0]); col++ {
			fmt.Printf("%v ", table[row][col])
		}
		fmt.Println()
	}
	fmt.Printf("\n")
}
