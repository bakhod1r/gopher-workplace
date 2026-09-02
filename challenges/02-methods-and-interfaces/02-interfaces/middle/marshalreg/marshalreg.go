// Package marshalreg — Gopher Workplace challenge.
package marshalreg

import (
	"errors"
	"strings"
)

// ErrNoCodec reports that nothing could encode the message.
var ErrNoCodec = errors.New("no codec")

// Codec encodes a message.
type Codec interface {
	Encode(fields []string) string
}

// CSVCodec joins with commas.
type CSVCodec struct{}

// Encode joins the fields with ",".
func (CSVCodec) Encode(fields []string) string {
	// TODO(candidate): comma-separated.
	panic("not implemented")
}

// PipeCodec joins with pipes.
type PipeCodec struct{}

// Encode joins the fields with "|".
func (PipeCodec) Encode(fields []string) string {
	// TODO(candidate): pipe-separated.
	panic("not implemented")
}

// Registry resolves a codec by content type.
type Registry struct {
	codecs  map[string]Codec
	fallbck Codec
}

// NewRegistry returns an empty registry.
func NewRegistry() *Registry {
	return &Registry{codecs: make(map[string]Codec)}
}

// Register binds a codec to a content type.
func (r *Registry) Register(contentType string, c Codec) {
	// TODO(candidate): store it.
	panic("not implemented")
}

// SetDefault sets the fallback codec.
func (r *Registry) SetDefault(c Codec) {
	// TODO(candidate): store the fallback.
	panic("not implemented")
}

// Encode resolves a codec by content type, falling back to the default.
//
// Examples:
//
//	reg.Encode("csv", []string{"a", "b"}) => "a,b", nil
//	unknown type with no default          => "", ErrNoCodec
func (r *Registry) Encode(contentType string, fields []string) (string, error) {
	// TODO(candidate): lookup, then default, then ErrNoCodec.
	panic("not implemented")
}

var _ = strings.Join
