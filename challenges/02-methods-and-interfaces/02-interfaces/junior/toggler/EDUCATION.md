# Toggler

## Intuition

A contract can mix commands and queries. Because `Toggle` mutates, both methods take pointer receivers so `*Switch` — and only `*Switch` — satisfies `Toggler`.

## Approach

1. `Toggle` sets `s.On = !s.On`.
2. `State` returns `s.On`.
3. `ToggleAll` toggles first, then counts the elements whose `State()` is true.

## Solution

```go
func (s *Switch) Toggle() { s.On = !s.On }

func (s *Switch) State() bool { return s.On }

func ToggleAll(ts []Toggler) int {
	n := 0
	for _, t := range ts {
		t.Toggle()
		if t.State() {
			n++
		}
	}
	return n
}
```

## Walkthrough

`{a: off, b: on}` after toggling becomes `{a: on, b: off}` — exactly one is on, so the count is 1.

## Pitfalls

- Counting before toggling, which reports the old state.
- A value receiver on `Toggle`, so the flip is lost.
- Mixing receivers (value `State`, pointer `Toggle`) — legal, but then only `*Switch` implements the interface anyway, so keep them consistent.
