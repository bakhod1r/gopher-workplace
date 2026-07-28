// Package tiers sets reward thresholds. A planted iota bug zeroes the first tier.
package tiers

// Tier is a reward threshold.
type Tier int

const (
	// CHANGE CODE BELOW THIS LINE
	Bronze Tier = iota * 100
	// CHANGE CODE ABOVE THIS LINE
	Silver
	Gold
)
