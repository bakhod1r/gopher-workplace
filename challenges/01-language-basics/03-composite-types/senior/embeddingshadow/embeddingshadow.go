// Package embeddingshadow uses struct embedding. A planted bug reads a shadowing
// outer field instead of the embedded one.
package embeddingshadow

// Base carries an ID.
type Base struct {
	ID int
}

// Entity embeds Base but also (mistakenly) declares its own ID that shadows it.
type Entity struct {
	Base
	ID int // shadows Base.ID
}

// BaseID must return the embedded Base's ID (set via the Base field), not the
// outer shadow.
func BaseID(e Entity) int {
	// CHANGE CODE BELOW THIS LINE
	return e.ID
	// CHANGE CODE ABOVE THIS LINE
}
