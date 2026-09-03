package hotaggregate

import (
	"bytes"
	"errors"
	"testing"
)

var (
	sinkT int64
	sinkC int
)

func TestAggregate(t *testing.T) {
	in := [][]byte{[]byte("1,2"), []byte("3"), []byte("-4,+5")}
	total, count, err := Aggregate(in, ',')
	if err != nil || total != 7 || count != 5 {
		t.Errorf("Aggregate = %d, %d, %v, want 7, 5, nil", total, count, err)
	}
}

func TestAggregateEmptyInputs(t *testing.T) {
	if total, count, err := Aggregate(nil, ','); err != nil || total != 0 || count != 0 {
		t.Errorf("Aggregate = %d, %d, %v, want 0, 0, nil", total, count, err)
	}
	if total, count, err := Aggregate([][]byte{nil, {}}, ','); err != nil || total != 0 || count != 0 {
		t.Errorf("Aggregate = %d, %d, %v, want 0, 0, nil", total, count, err)
	}
}

func TestAggregateSyntaxErrors(t *testing.T) {
	for _, in := range []string{"1,,2", "1,x", "-", "+", "a"} {
		if _, _, err := Aggregate([][]byte{[]byte(in)}, ','); !errors.Is(err, ErrSyntax) {
			t.Errorf("Aggregate(%q) = %v, want ErrSyntax", in, err)
		}
	}
}

func TestAggregateAllocatesNothing(t *testing.T) {
	lines := make([][]byte, 64)
	for i := range lines {
		lines[i] = bytes.Repeat([]byte("12345,"), 16)
		lines[i] = lines[i][:len(lines[i])-1]
	}
	n := testing.AllocsPerRun(50, func() { sinkT, sinkC, _ = Aggregate(lines, ',') })
	if n != 0 {
		t.Errorf("Aggregate made %v allocations, want 0", n)
	}
}

func TestAggregateErrorPathAllocatesNothing(t *testing.T) {
	lines := [][]byte{[]byte("1,x")}
	var err error
	n := testing.AllocsPerRun(50, func() { _, _, err = Aggregate(lines, ',') })
	_ = err
	if n != 0 {
		t.Errorf("the error path made %v allocations, want 0: return the sentinel", n)
	}
}

func BenchmarkAggregate(b *testing.B) {
	lines := make([][]byte, 256)
	for i := range lines {
		lines[i] = bytes.Repeat([]byte("12345,"), 16)
		lines[i] = lines[i][:len(lines[i])-1]
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkT, sinkC, _ = Aggregate(lines, ',')
	}
}
