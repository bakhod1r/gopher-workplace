// Package counternilmap tallies word counts. A planted bug returns an unmade
// (nil) map from the constructor, so the first write panics.
package counternilmap

// Tally counts occurrences of each string in words.
func Tally(words []string) map[string]int {
	// CHANGE CODE BELOW THIS LINE
	var m map[string]int
	// CHANGE CODE ABOVE THIS LINE
	for _, w := range words {
		m[w]++
	}
	return m
}
