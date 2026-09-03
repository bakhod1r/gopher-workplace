package encodeexact

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

var sink []byte

func TestEncode(t *testing.T) {
	got := Encode([]Rec{{ID: 1, Name: "a"}, {ID: 22, Name: "bb"}})
	if !bytes.Equal(got, []byte("1:a\n22:bb")) {
		t.Errorf("Encode = %q, want \"1:a\\n22:bb\"", got)
	}
}

func TestEncodeEmpty(t *testing.T) {
	if got := Encode(nil); len(got) != 0 {
		t.Errorf("Encode = %q, want empty", got)
	}
	if got := Encode([]Rec{}); len(got) != 0 {
		t.Errorf("Encode = %q, want empty", got)
	}
}

func TestEncodeEdges(t *testing.T) {
	got := Encode([]Rec{{ID: -5, Name: ""}})
	if !bytes.Equal(got, []byte("-5:")) {
		t.Errorf("Encode = %q, want \"-5:\"", got)
	}
}

func TestEncodeLarge(t *testing.T) {
	recs := make([]Rec, 500)
	var want strings.Builder
	for i := range recs {
		recs[i] = Rec{ID: i * 37, Name: strings.Repeat("n", i%5)}
		if i > 0 {
			want.WriteByte('\n')
		}
		want.WriteString(strconv.Itoa(recs[i].ID))
		want.WriteByte(':')
		want.WriteString(recs[i].Name)
	}
	if got := Encode(recs); string(got) != want.String() {
		t.Errorf("Encode produced %d bytes, want %d", len(got), want.Len())
	}
}

func TestEncodeAllocatesOnce(t *testing.T) {
	recs := make([]Rec, 64)
	for i := range recs {
		recs[i] = Rec{ID: i, Name: "name"}
	}
	n := testing.AllocsPerRun(50, func() { sink = Encode(recs) })
	if n > 1 {
		t.Errorf("Encode made %v allocations, want 1: compute the length first", n)
	}
}
