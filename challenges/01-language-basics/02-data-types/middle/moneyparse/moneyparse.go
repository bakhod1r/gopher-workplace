// Package moneyparse parses a dollar string into integer cents.
package moneyparse

// Cents parses a string like "12.34" or "0.05" or "7" into the number of cents
// (1234, 5, 700). Returns (0,false) on bad format or more than two decimals.
//
// TODO(candidate): implement by scanning; keep exact integer cents, no float.
func Cents(s string) (int, bool) {
	panic("not implemented")
}
