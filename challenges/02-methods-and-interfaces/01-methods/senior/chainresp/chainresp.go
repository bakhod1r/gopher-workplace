// Package chainresp — Gopher Workplace challenge.
package chainresp

// Handler is an interface for chain links.
type Handler interface {
	Handle(req int) string
	SetNext(Handler)
}

// BaseHandler implements common chaining.
type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(next Handler) { b.next = next }
func (b *BaseHandler) Next(req int) string {
	if b.next != nil {
		return b.next.Handle(req)
	}
	return "unhandled"
}

// H10 handles 10.
type H10 struct{ BaseHandler }

func (h *H10) Handle(req int) string {
	if req == 10 {
		return "ten"
	}
	return h.Next(req)
}

// H20 handles 20.
type H20 struct{ BaseHandler }

func (h *H20) Handle(req int) string {
	// TODO(candidate): if req == 20 return "twenty", else h.Next(req)
	panic("not implemented")
}
