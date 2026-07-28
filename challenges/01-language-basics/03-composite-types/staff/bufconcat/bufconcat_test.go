package bufconcat

import (
	"bytes"
	"testing"
)

func TestConcat(t *testing.T) {
	got := Concat([]byte("foo"), []byte("bar"))
	if !bytes.Equal(got, []byte("foobar")) {
		t.Errorf("Concat=%q; want foobar", got)
	}
	got = Concat(nil, []byte("x"))
	if !bytes.Equal(got, []byte("x")) {
		t.Errorf("Concat nil=%q; want x", got)
	}
}
