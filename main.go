package main

import (
	"fmt"
	"name-clash/names"
)

func main() {
	nameNodes, _ := names.CreateNameGraph("resources/test.txt", "Test", 1)
	tournament := names.CreateTournament(nameNodes)
	var champ *names.Name = &names.Name{Text: "None"}
	fmt.Printf("Tournament: %v\nChampion: %v\n\n", tournament, champ.Text)
	tournament, champ = names.Vote(tournament, 0, 0)
	fmt.Printf("Tournament: %v\nChampion: %v\n\n", tournament, champ.Text)
	tournament, champ = names.Vote(tournament, 0, 0)
	fmt.Printf("Tournament: %v\nChampion: %v\n\n", tournament, champ.Text)
	tournament, champ = names.Vote(tournament, 0, 0)
	fmt.Printf("Tournament: %v\nChampion: %v\n\n", tournament, champ.Text)
	tournament, champ = names.Vote(tournament, 0, 0)
	fmt.Printf("Tournament: %v\nChampion: %v\n\n", tournament, champ.Text)
	tournament, champ = names.Vote(tournament, 0, 0)
	fmt.Printf("Tournament: %v\nChampion: %v\n\n", tournament, champ.Text)
	tournament, champ = names.Vote(tournament, 0, 0)
	fmt.Printf("Tournament: %v\nChampion: %v\n\n", tournament, champ.Text)
	tournament, champ = names.Vote(tournament, 0, 0)
	fmt.Printf("Tournament: %v\nChampion: %v\n\n", tournament, champ.Text)
	//graph.ExportGraphML(nameNodes, nameEdges)
}
