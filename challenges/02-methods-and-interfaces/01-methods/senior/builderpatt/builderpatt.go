// Package builderpatt — Gopher Workplace challenge.
package builderpatt

// Request holds HTTP request params.
type Request struct {
	Method string
	URL    string
	Auth   string
}

// RequestBuilder builds a Request fluently.
type RequestBuilder struct {
	req Request
}

// NewBuilder starts building.
func NewBuilder() *RequestBuilder {
	return &RequestBuilder{}
}

// Method sets the HTTP method.
func (b *RequestBuilder) Method(m string) *RequestBuilder {
	b.req.Method = m
	return b
}

// URL sets the URL.
func (b *RequestBuilder) URL(u string) *RequestBuilder {
	// TODO(candidate): set req.URL and return b
	panic("not implemented")
}

// Auth sets the auth token.
func (b *RequestBuilder) Auth(t string) *RequestBuilder {
	// TODO(candidate): set req.Auth and return b
	panic("not implemented")
}

// Build returns the final Request.
func (b *RequestBuilder) Build() Request {
	return b.req
}
