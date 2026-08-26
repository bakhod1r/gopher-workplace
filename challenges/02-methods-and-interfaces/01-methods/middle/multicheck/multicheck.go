// Package multicheck — Gopher Workplace challenge.
package multicheck

import "strings"

// User holds user data.
type User struct {
	Name  string
	Email string
	Age   int
}

// Validate checks all fields and returns a slice of error strings.
// Returns nil if all checks pass.
//
// Rules:
//   - Name must not be empty.
//   - Email must contain "@".
//   - Age must be >= 0.
//
// Examples:
//
//	User{"", "bad", -1}.Validate() => ["name is required", "invalid email", "age must be non-negative"]
//	User{"Go", "go@go.dev", 10}.Validate() => nil
func (u User) Validate() []string {
	// TODO(candidate): check each rule, collect errors.
	_ = strings.Contains // hint
	panic("not implemented")
}
