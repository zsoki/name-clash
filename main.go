package main

import (
	"name-clash/graph"
	"name-clash/names"
)

func main() {
	nameNodes, nameEdges := names.CreateNameGraph("resources/noi.txt", "Anna", 1)
	matchups := names.CreateRandomMatchups(nameNodes)
	println(matchups)
	graph.ExportGraphML(nameNodes, nameEdges)
}
