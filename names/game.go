package names

import "math/rand"

type Tournament struct {
	Remaining [][2]*Name
	Winners   []*Name
	Losers    []*Name
}

func CreateTournament(names []*Name) Tournament {
	// TODO what to do if not even?
	if len(names)%2 != 0 {
		names = names[1:]
	}

	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })

	matches := make([][2]*Name, len(names)/2)
	for i := 0; i < len(names); i += 2 {
		matches[i/2] = [2]*Name{names[i], names[i+1]}
	}

	return Tournament{Remaining: matches}
}

func Vote(tournament Tournament, matchIdx, winnerIdx int) Tournament {
	loserIdx := 1 - winnerIdx
	winner := tournament.Remaining[matchIdx][winnerIdx]
	loser := tournament.Remaining[matchIdx][loserIdx]

	winner.Wins += 1
	loser.Wins -= 1

	tournament.Winners = append(tournament.Winners, winner)
	tournament.Losers = append(tournament.Losers, loser)

	// TODO handle winners/losers bracket?

	tournament.Remaining = append(tournament.Remaining[:matchIdx], tournament.Remaining[matchIdx+1:]...)
	return tournament
}
