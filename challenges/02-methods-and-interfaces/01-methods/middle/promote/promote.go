// Package promote — Gopher Workplace challenge.
package promote

// Base provides a Name.
type Base struct {
	Name string
}

// Hello returns "Hello from <Name>".
func (b Base) Hello() string {
	return "Hello from " + b.Name
}

// Extended embeds Base and adds an Extra field.
type Extended struct {
	Base
	Extra string
}

// Describe returns "<Hello()> [<Extra>]".
// It should USE the promoted Hello method from Base.
//
// Examples:
//
//	Extended{Base{"Go"}, "fast"}.Describe() => "Hello from Go [fast]"
func (e Extended) Describe() string {
	// TODO(candidate): use e.Hello() (promoted) and e.Extra.
	panic("not implemented")
}
