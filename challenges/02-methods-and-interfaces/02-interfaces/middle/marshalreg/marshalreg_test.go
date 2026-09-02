package marshalreg

import (
	"errors"
	"testing"
)

func TestCodecs(t *testing.T) {
	if got := (CSVCodec{}).Encode([]string{"a", "b"}); got != "a,b" {
		t.Errorf("CSV = %q, want \"a,b\"", got)
	}
	if got := (PipeCodec{}).Encode([]string{"a", "b"}); got != "a|b" {
		t.Errorf("Pipe = %q, want \"a|b\"", got)
	}
	if got := (CSVCodec{}).Encode(nil); got != "" {
		t.Errorf("CSV(nil) = %q, want empty", got)
	}
}

func TestEncodeRegistered(t *testing.T) {
	r := NewRegistry()
	r.Register("csv", CSVCodec{})

	got, err := r.Encode("csv", []string{"a", "b"})
	if err != nil || got != "a,b" {
		t.Errorf("Encode = %q, %v", got, err)
	}
}

func TestEncodeFallback(t *testing.T) {
	r := NewRegistry()
	r.Register("csv", CSVCodec{})
	r.SetDefault(PipeCodec{})

	got, err := r.Encode("unknown", []string{"a", "b"})
	if err != nil || got != "a|b" {
		t.Errorf("Encode = %q, %v; want \"a|b\", nil", got, err)
	}
}

func TestEncodeNoCodec(t *testing.T) {
	r := NewRegistry()
	got, err := r.Encode("unknown", []string{"a"})
	if !errors.Is(err, ErrNoCodec) {
		t.Errorf("err = %v, want ErrNoCodec", err)
	}
	if got != "" {
		t.Errorf("got = %q, want empty", got)
	}
}

func TestRegisterOverrides(t *testing.T) {
	r := NewRegistry()
	r.Register("x", CSVCodec{})
	r.Register("x", PipeCodec{})
	if got, _ := r.Encode("x", []string{"a", "b"}); got != "a|b" {
		t.Errorf("Encode = %q, want \"a|b\"", got)
	}
}
