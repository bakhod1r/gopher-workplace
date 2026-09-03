package concatloop

import (
	"strings"
	"testing"
)

var sink string

func TestJoin(t *testing.T) {
	if got := Join([]string{"a", "bc", "d"}); got != "abcd" {
		t.Errorf("Join = %q, want \"abcd\"", got)
	}
	if got := Join(nil); got != "" {
		t.Errorf("Join = %q, want empty", got)
	}
	if got := Join([]string{"", ""}); got != "" {
		t.Errorf("Join = %q, want empty", got)
	}
	if got := Join([]string{"solo"}); got != "solo" {
		t.Errorf("Join = %q, want \"solo\"", got)
	}
}

func TestJoinLong(t *testing.T) {
	parts := make([]string, 200)
	for i := range parts {
		parts[i] = "chunk"
	}
	got := Join(parts)
	if got != strings.Repeat("chunk", 200) {
		t.Errorf("Join produced %d bytes, want %d", len(got), 200*5)
	}
}

func TestJoinAllocationsAreBounded(t *testing.T) {
	parts := make([]string, 128)
	for i := range parts {
		parts[i] = "chunk"
	}
	n := testing.AllocsPerRun(20, func() { sink = Join(parts) })
	if n > 3 {
		t.Errorf("Join made %v allocations for 128 parts, want at most 3: build in one buffer", n)
	}
}
