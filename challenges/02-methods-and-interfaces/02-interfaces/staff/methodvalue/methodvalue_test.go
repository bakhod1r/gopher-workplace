package methodvalue

import "testing"

func TestCounterGetSet(t *testing.T) {
	c := &Counter{N: 1}
	if got := c.Get(); got != 1 {
		t.Errorf("Get = %d, want 1", got)
	}
	c.Set(5)
	if got := c.Get(); got != 5 {
		t.Errorf("Get = %d, want 5", got)
	}
}

func TestValCounterGet(t *testing.T) {
	if got := (ValCounter{N: 3}).Get(); got != 3 {
		t.Errorf("Get = %d, want 3", got)
	}
}

func TestBindValueCapturesACopy(t *testing.T) {
	v := ValCounter{N: 1}
	f := BindValue(v)

	v.N = 99 // the closure holds a copy made at bind time

	if got := f(); got != 1 {
		t.Errorf("f() = %d, want 1: a value-receiver method value captures a copy", got)
	}
}

func TestBindPointerSeesMutations(t *testing.T) {
	c := &Counter{N: 1}
	f := BindPointer(c)

	c.Set(99)

	if got := f(); got != 99 {
		t.Errorf("f() = %d, want 99: a pointer-receiver method value follows the pointer", got)
	}
}

func TestMethodExpression(t *testing.T) {
	get := GetExpr()
	if got := get(ValCounter{N: 7}); got != 7 {
		t.Errorf("get = %d, want 7", got)
	}
	if got := get(ValCounter{N: 0}); got != 0 {
		t.Errorf("get = %d, want 0", got)
	}
}

func TestBoundValueSurvivesRebinding(t *testing.T) {
	v := ValCounter{N: 1}
	first := BindValue(v)
	v.N = 2
	second := BindValue(v)

	if first() != 1 {
		t.Errorf("first() = %d, want 1", first())
	}
	if second() != 2 {
		t.Errorf("second() = %d, want 2", second())
	}
}

func TestPointerBindingIsShared(t *testing.T) {
	c := &Counter{N: 1}
	a := BindPointer(c)
	b := BindPointer(c)

	c.Set(42)
	if a() != 42 || b() != 42 {
		t.Errorf("a = %d, b = %d; both should see 42", a(), b())
	}
}

func TestSatisfiesGetter(t *testing.T) {
	var g Getter = &Counter{N: 4}
	if g.Get() != 4 {
		t.Errorf("Get = %d, want 4", g.Get())
	}

	var g2 Getter = ValCounter{N: 5}
	if g2.Get() != 5 {
		t.Errorf("Get = %d, want 5", g2.Get())
	}
}

func TestCallbackRegistrationPitfall(t *testing.T) {
	// The bug this puzzle is about: registering a value-receiver method value
	// as a callback freezes the state at registration time.
	v := ValCounter{N: 0}
	callbacks := []func() int{}
	for i := 1; i <= 3; i++ {
		v.N = i
		callbacks = append(callbacks, BindValue(v))
	}

	want := []int{1, 2, 3}
	for i, cb := range callbacks {
		if got := cb(); got != want[i] {
			t.Errorf("callbacks[%d] = %d, want %d", i, got, want[i])
		}
	}
}
