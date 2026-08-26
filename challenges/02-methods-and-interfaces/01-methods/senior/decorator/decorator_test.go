package decorator

import "testing"

func TestDecorator(t *testing.T) {
	c := &Component{}
	d := &Decorator{Comp: c}

	if got := d.DoWork(); got != "[work]" {
		t.Errorf("DoWork() = %q, want %q", got, "[work]")
	}
}
