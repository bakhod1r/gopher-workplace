// Package typeswitchtab — Gopher Workplace challenge.
package typeswitchtab

import "reflect"

// Decoder labels a value.
type Decoder interface {
	Decode(v any) string
}

// DecodeSwitch classifies v with a type switch.
//
// Examples:
//
//	DecodeSwitch(1)    => "int"
//	DecodeSwitch(3.5)  => "unknown"
func DecodeSwitch(v any) string {
	// TODO(candidate): type switch over int, string, bool.
	panic("not implemented")
}

// Table dispatches by dynamic type through a map.
type Table struct {
	labels map[reflect.Type]string
}

// NewTable returns an empty table.
func NewTable() *Table {
	return &Table{labels: make(map[reflect.Type]string)}
}

// Register binds a label to the dynamic type of sample.
func (t *Table) Register(sample any, label string) {
	// TODO(candidate): key by reflect.TypeOf(sample).
	panic("not implemented")
}

// DecodeTable classifies v via the table, returning "unknown" when the type
// is not registered.
func (t *Table) DecodeTable(v any) string {
	// TODO(candidate): nil-safe reflect lookup.
	panic("not implemented")
}
