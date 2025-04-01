package names

import (
	"math/rand"
)

type Tournament struct {
	Remaining [][2]*Name
	Winners   []*Name
	Losers    []*Name
}

func (tournament Tournament) String() string {
	returnString := "\nRemaining:"
	for _, n := range tournament.Remaining {
		returnString += "\n\t" + n[0].Text + " " + n[1].Text
	}
	returnString += "\nWinners:"
	for _, n2 := range tournament.Winners {
		returnString += "\n\t" + n2.Text
	}
	returnString += "\nLosers:"
	for _, n3 := range tournament.Losers {
		returnString += "\n\t" + n3.Text
	}
	return returnString
}

func CreateTournament(names []*Name) Tournament {
	// TODO what to do if not even?
	if len(names)%2 != 0 {
		names = names[1:]
	}

	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })

	matches := createPairs(names)

	return Tournament{Remaining: matches}
}

// Vote returns a modified Tournament after the vote and the champion of the tournament, if any.
func Vote(tournament Tournament, matchIdx, winnerIdx int) (Tournament, *Name) {
	if len(tournament.Remaining) == 0 {
		return tournament, &Name{Text: "There's nothing to vote"}
	}

	loserIdx := 1 - winnerIdx
	winner := tournament.Remaining[matchIdx][winnerIdx]
	loser := tournament.Remaining[matchIdx][loserIdx]

	winner.Wins += 1
	loser.Losses += 1

	if winner.Losses == 0 { // Winner stays in original bracket.
		tournament.Winners = append(tournament.Winners, winner)
	} else {
		tournament.Losers = append(tournament.Losers, winner)
	}

	if loser.Losses == 1 { // Move to loser's bracket after first loss. 2 loss is eliminated.
		tournament.Losers = append(tournament.Losers, loser)
	}

	tournament.Remaining = append(tournament.Remaining[:matchIdx], tournament.Remaining[matchIdx+1:]...)
	var champion *Name = &Name{Text: "None"}

	if len(tournament.Remaining) == 0 {
		var matches [][2]*Name

		if len(tournament.Winners) == 1 && len(tournament.Losers) == 1 {
			// Finals happening
			matches = append(make([][2]*Name, 0), [2]*Name{tournament.Winners[0], tournament.Losers[0]})
			tournament.Winners = nil
		} else if len(tournament.Winners) == 1 && len(tournament.Losers) == 0 {
			// Champion can be announced
			champion = tournament.Winners[0]
			tournament.Winners = nil
		} else {
			// Create new matchups from winners' and losers' bracket
			if len(tournament.Winners) > 1 {
				matches = createPairs(tournament.Winners)
				tournament.Winners = nil
			}
			matches = append(matches, createPairs(tournament.Losers)...)
		}

		tournament.Losers = nil
		tournament.Remaining = matches
	}

	return tournament, champion
}

func createPairs(names []*Name) [][2]*Name {
	matches := make([][2]*Name, len(names)/2)
	for i := 0; i < len(names); i += 2 {
		matches[i/2] = [2]*Name{names[i], names[i+1]}
	}
	return matches
}
