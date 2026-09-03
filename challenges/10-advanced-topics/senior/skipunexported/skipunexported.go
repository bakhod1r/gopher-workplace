// Package skipunexported — Gopher Workplace challenge.
package skipunexported

import "reflect"

// SumInts returns the total of v's exported int fields.
//
// Unexported fields can be read as reflect Values but not converted back
// through Interface, and reaching for them panics.
//
// Examples:
//
//	SumInts(mix{A: 1, b: 2}) => 1
func SumInts(v any) int64 {
	// CHANGE CODE BELOW THIS LINE
	rv := reflect.ValueOf(v)
	if !rv.IsValid() || rv.Kind() != reflect.Struct {
		return 0
	}
	var total int64
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		if f.Kind() != reflect.Int {
			continue
		}
		total += f.Interface().(int64)
	}
	return total
	// CHANGE CODE ABOVE THIS LINE
}
