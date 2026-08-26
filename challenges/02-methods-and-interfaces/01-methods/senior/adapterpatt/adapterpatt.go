// Package adapterpatt — Gopher Workplace challenge.
package adapterpatt

import "strconv"

// LegacySystem provides data as strings.
type LegacySystem struct{}

func (LegacySystem) GetStringData() string { return "123" }

// ModernAdapter adapts LegacySystem to return ints.
type ModernAdapter struct {
	legacy LegacySystem
}

// GetIntData returns the string data parsed as an int. If it fails, return 0.
func (a *ModernAdapter) GetIntData() int {
	// TODO(candidate): call legacy.GetStringData(), use strconv.Atoi
	_ = strconv.Atoi
	panic("not implemented")
}
