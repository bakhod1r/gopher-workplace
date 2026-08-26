package bloomfilter

import "testing"

func TestBloom(t *testing.T) {
	f := &Filter{}
	f.Add("hello")

	if !f.MightContain("hello") {
		t.Error("should contain hello")
	}
	if f.MightContain("world") {
		t.Error("should not contain world")
	}
	if !f.MightContain("ho") {
		t.Log("false positive, which is fine in Bloom")
	}
}
