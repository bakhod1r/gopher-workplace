// Package sortslice — Gopher Workplace challenge.
package sortslice

// Player is one leaderboard entry.
type Player struct {
	Name  string
	Score int
}

// ByScore orders players by descending score, then ascending name.
type ByScore []Player

// Len returns the number of players.
func (b ByScore) Len() int {
	// TODO(candidate): length.
	panic("not implemented")
}

// Less reports whether i sorts before j.
//
// Higher scores come first; equal scores are ordered by name.
func (b ByScore) Less(i, j int) bool {
	// TODO(candidate): descending score, then ascending name.
	panic("not implemented")
}

// Swap exchanges two players.
func (b ByScore) Swap(i, j int) {
	// TODO(candidate): swap in place.
	panic("not implemented")
}

// TopN sorts players and returns the first n names.
func TopN(players []Player, n int) []string {
	// TODO(candidate): sort.Sort(ByScore(players)), then take n names.
	panic("not implemented")
}
