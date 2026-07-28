// Package tiers ranks scores using an iota expression.
package tiers

// Tier is a reward tier threshold.
type Tier int

// Bronze=100, Silver=200, Gold=300 via a repeated iota expression.
//
// TODO(candidate): define with (iota+1)*100 written once.
const (
	Bronze Tier = 0
	Silver Tier = 0
	Gold   Tier = 0
)

// Rank returns the highest tier whose threshold is <= score, or 0 below Bronze.
//
// TODO(candidate): implement.
func Rank(score int) Tier {
	panic("not implemented")
}
