package poolalias

import (
	"bytes"
	"strconv"
	"testing"
)

func TestEncode(t *testing.T) {
	if got := Encode([]int{1, 2, 3}); !bytes.Equal(got, []byte("1,2,3")) {
		t.Errorf("Encode = %q, want \"1,2,3\"", got)
	}
	if got := Encode(nil); len(got) != 0 {
		t.Errorf("Encode = %q, want empty", got)
	}
}

func TestEarlierResultsSurviveLaterCalls(t *testing.T) {
	first := Encode([]int{111, 222})
	for i := 0; i < 50; i++ {
		Encode([]int{999, 888})
	}
	if !bytes.Equal(first, []byte("111,222")) {
		t.Errorf("first = %q, want \"111,222\": the result was a view of a pooled buffer", first)
	}
}

func TestResultsAreIndependentOfEachOther(t *testing.T) {
	got := make([][]byte, 0, 32)
	for i := 0; i < 32; i++ {
		got = append(got, Encode([]int{i, i * 2}))
	}
	for i, b := range got {
		want := Reference(i)
		if !bytes.Equal(b, want) {
			t.Fatalf("result %d = %q, want %q", i, b, want)
		}
	}
}

// Reference renders the expected output without the pool.
func Reference(i int) []byte {
	var out []byte
	out = strconv.AppendInt(out, int64(i), 10)
	out = append(out, ',')
	return strconv.AppendInt(out, int64(i*2), 10)
}
