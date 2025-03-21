package names

import "math/rand"

func CreateRandomMatchups(names []*Name) [][2]*Name {
	// TODO what to do if not even?
	if len(names)%2 != 0 {
		names = names[1:]
	}

	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })

	matchups := make([][2]*Name, len(names)/2)
	for i := 0; i < len(names); i += 2 {
		matchups[i/2] = [2]*Name{names[i], names[i+1]}
	}
	return matchups
}
