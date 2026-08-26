// Package withopt — Gopher Workplace challenge.
package withopt

// Server holds configuration.
type Server struct {
	Host    string
	Port    int
	Timeout int // seconds
}

// Option is a function that configures a Server.
type Option func(*Server)

// WithHost returns an Option that sets the host.
func WithHost(host string) Option {
	// TODO(candidate): return a function that sets s.Host.
	panic("not implemented")
}

// WithPort returns an Option that sets the port.
func WithPort(port int) Option {
	// TODO(candidate): return a function that sets s.Port.
	panic("not implemented")
}

// WithTimeout returns an Option that sets the timeout.
func WithTimeout(timeout int) Option {
	// TODO(candidate): return a function that sets s.Timeout.
	panic("not implemented")
}

// NewServer creates a Server with defaults, then applies options.
//
// Defaults: Host="localhost", Port=8080, Timeout=30.
//
// Examples:
//
//	NewServer()                          => {localhost, 8080, 30}
//	NewServer(WithPort(9090))            => {localhost, 9090, 30}
//	NewServer(WithHost("0.0.0.0"), WithPort(3000)) => {0.0.0.0, 3000, 30}
func NewServer(opts ...Option) Server {
	// TODO(candidate): create default server, apply each option, return.
	panic("not implemented")
}
