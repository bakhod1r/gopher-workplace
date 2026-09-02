// Package assertsafe — Gopher Workplace challenge.
package assertsafe

// Handler consumes a payload if it understands it.
type Handler interface {
	Handle(payload any) bool
}

// IntHandler sums int payloads.
type IntHandler struct {
	Sum int
}

// Handle accepts only ints.
//
// Examples:
//
//	h := &IntHandler{}; h.Handle(3)   => true, h.Sum == 3
//	h.Handle("3")                     => false
func (h *IntHandler) Handle(payload any) bool {
	// TODO(candidate): accept ints only.
	panic("not implemented")
}

// TextHandler collects string payloads.
type TextHandler struct {
	Seen []string
}

// Handle accepts only strings.
func (h *TextHandler) Handle(payload any) bool {
	// TODO(candidate): accept strings only.
	panic("not implemented")
}

// Dispatch offers payload to every handler and counts the acceptances.
func Dispatch(hs []Handler, payload any) int {
	// TODO(candidate): offer to all, count true results.
	panic("not implemented")
}
