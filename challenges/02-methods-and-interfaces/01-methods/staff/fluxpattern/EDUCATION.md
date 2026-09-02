# Flux Store

## Intuition

Flux is the observer pattern with a rule attached: state is only ever written by
a reducer, and the reducer's only input is an action. That constraint is what
makes time-travel debugging possible — replay the action log and you replay the
state.

## Approach

1. Switch on the action name.
2. Apply the state change for the actions you know.
3. Ignore everything else.

## Solution

```go
func (s *Store) Dispatch(action string) {
	switch action {
	case "INC":
		s.Count++
	case "DEC":
		s.Count--
	}
}
```

## Walkthrough

The test dispatches `INC`, `INC`, `DEC` against a zero-valued store: 0 → 1 → 2 →
1. Each call is independent; the store carries the only state.

## Pitfalls

- **Value receiver.** Every dispatch mutates a copy and `Count` stays 0.
- **Panicking on an unknown action.** In a real store many reducers see every
  action and most ignore it; panicking would make composition impossible.
- **Exported `Count` with outside writers.** The test reads it directly for
  convenience, but any code that writes it bypasses the audit trail the pattern
  exists to provide.

## Where this grows

A real store also holds subscribers and notifies them after each dispatch — the
`observer` puzzle is the other half. Keeping the reducer pure (no I/O, no
clock, no randomness) is what keeps replay deterministic.
