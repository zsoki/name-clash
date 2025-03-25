package main

import (
	"fmt"
	"name-clash/graph"
	"name-clash/names"
)

func main() {
	nameNodes, nameEdges := names.CreateNameGraph("resources/noi.txt", "Anna", 1)
	tournament := names.CreateTournament(nameNodes)
	fmt.Printf("Tournament: %v\n", tournament)
	tournament = names.Vote(tournament, 0, 0)
	tournament = names.Vote(tournament, 0, 1)
	graph.ExportGraphML(nameNodes, nameEdges)
}
