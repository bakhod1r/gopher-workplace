// Package idvalidate — Gopher Workplace challenge.
package idvalidate

// UserID identifies a user.
type UserID string

// OrderID identifies an order.
type OrderID string

// ValidID reports whether v is non-empty and has the given prefix.
func ValidID[T ~string](v T, prefix string) bool {
	// TODO(candidate): check emptiness and the prefix.
	panic("not implemented")
}
