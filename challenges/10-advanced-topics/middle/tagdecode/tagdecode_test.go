package tagdecode

import (
	"errors"
	"testing"
)

type config struct {
	Host    string `env:"HOST"`
	Port    int    `env:"PORT"`
	Debug   bool   `env:"DEBUG"`
	Ignored string
	skipped string `env:"SKIPPED"`
	Dashed  string `env:"-"`
}

func TestDecode(t *testing.T) {
	var c config
	src := map[string]string{
		"HOST": "localhost", "PORT": "8080", "DEBUG": "true",
		"SKIPPED": "x", "-": "y", "UNKNOWN": "z",
	}
	if err := Decode(src, &c); err != nil {
		t.Fatalf("Decode = %v, want nil", err)
	}
	if c.Host != "localhost" || c.Port != 8080 || !c.Debug {
		t.Errorf("c = %+v, want {localhost 8080 true ...}", c)
	}
	if c.Ignored != "" || c.skipped != "" || c.Dashed != "" {
		t.Errorf("c = %+v: untagged, unexported and \"-\" fields must be left alone", c)
	}
}

func TestDecodeLeavesMissingKeys(t *testing.T) {
	c := config{Host: "keep", Port: 1}
	if err := Decode(map[string]string{"DEBUG": "1"}, &c); err != nil {
		t.Fatal(err)
	}
	if c.Host != "keep" || c.Port != 1 || !c.Debug {
		t.Errorf("c = %+v, want the untouched fields preserved", c)
	}
}

func TestDecodeBadValues(t *testing.T) {
	var c config
	if err := Decode(map[string]string{"PORT": "eighty"}, &c); err == nil {
		t.Error("want an error for a non-numeric port, got nil")
	}
	if err := Decode(map[string]string{"DEBUG": "maybe"}, &c); err == nil {
		t.Error("want an error for a non-boolean debug, got nil")
	}
}

func TestDecodeBadTarget(t *testing.T) {
	for _, dst := range []any{config{}, nil, (*config)(nil), new(int)} {
		if err := Decode(nil, dst); !errors.Is(err, ErrTarget) {
			t.Errorf("Decode(%#v) = %v, want ErrTarget", dst, err)
		}
	}
}

func TestDecodeUnsupportedKind(t *testing.T) {
	var bad struct {
		F float64 `env:"F"`
	}
	if err := Decode(map[string]string{"F": "1.5"}, &bad); err == nil {
		t.Error("want an error for an unsupported kind, got nil")
	}
}
