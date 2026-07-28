// Package jsontag serializes a user to JSON. A planted bug has the wrong tag, so
// the field name is wrong in output.
package jsontag

import "encoding/json"

// User is serialized to JSON with snake_case keys.
type User struct {
	FirstName string `json:"first_name"`
	// CHANGE CODE BELOW THIS LINE
	LastName string `json:"lastName"`
	// CHANGE CODE ABOVE THIS LINE
}

// Marshal returns the JSON encoding of u.
func Marshal(u User) (string, error) {
	b, err := json.Marshal(u)
	return string(b), err
}
