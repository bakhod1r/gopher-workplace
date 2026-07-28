package bytecat

import (
	"bytes"
	"testing"
)

func TestFromBytes(t *testing.T) {
	if got := FromBytes([]byte{'G', 'o'}); got != "Go" {
		t.Errorf("FromBytes=%q; want Go", got)
	}
	if got := FromBytes(nil); got != "" {
		t.Errorf("FromBytes(nil)=%q; want empty", got)
	}
}

func TestToBytes(t *testing.T) {
	if got := ToBytes("Go"); !bytes.Equal(got, []byte{'G', 'o'}) {
		t.Errorf("ToBytes=%v; want [71 111]", got)
	}
}
