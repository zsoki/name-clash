package main

import (
	"name-clash/graph"
	"name-clash/names"
)

func main() {
	names, edges := names.Run("resources/noi.txt", "Anna", 1)
	graph.ExportGraphML(names, edges)
}
