// Package authcheck — Gopher Workplace challenge.
package authcheck

// Authenticator reports whether a credential grants access.
type Authenticator interface {
	Allow() bool
}

// Token is a bearer credential.
type Token struct {
	Value string
}

// Allow reports whether the token grants access.
//
// Examples:
//
//	Token{Value: "abc"}.Allow()     => true
//	Token{Value: "expired"}.Allow() => false
func (t Token) Allow() bool {
	// TODO(candidate): non-empty and not "expired".
	panic("not implemented")
}

// Guest is an anonymous visitor.
type Guest struct{}

// Allow reports whether the guest grants access.
func (g Guest) Allow() bool {
	// TODO(candidate): a guest never gets in.
	panic("not implemented")
}

// CanEnter reports whether a is allowed through the gate.
func CanEnter(a Authenticator) bool {
	// TODO(candidate): ask the authenticator.
	panic("not implemented")
}
