// Package describer — Gopher Workplace challenge.
package describer

// Describer renders itself as one line of text.
type Describer interface {
	Describe() string
}

// User is an account.
type User struct {
	Name string
}

// Describe renders the user.
//
// Examples:
//
//	User{Name: "ann"}.Describe() => "user ann"
func (u User) Describe() string {
	// TODO(candidate): "user <Name>".
	panic("not implemented")
}

// Server is a host/port pair.
type Server struct {
	Host string
	Port int
}

// Describe renders the server.
//
// Examples:
//
//	Server{Host: "db", Port: 5432}.Describe() => "server db:5432"
func (s Server) Describe() string {
	// TODO(candidate): "server <Host>:<Port>".
	panic("not implemented")
}

// DescribeAll renders every element, in order.
func DescribeAll(ds []Describer) []string {
	// TODO(candidate): collect Describe() for each element.
	panic("not implemented")
}
