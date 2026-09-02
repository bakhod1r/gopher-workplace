# Observer Pattern

## Intuition

Observer inverts a dependency: instead of the interested code polling the
subject, the subject pushes. Because Go treats functions as values, the
"observer interface" collapses into a single func type — registration is just
`append`.

## Approach

1. Commit the new state.
2. Fan out to every registered observer with that value.

## Solution

```go
func (s *Subject) SetState(val int) {
	s.state = val
	for _, o := range s.observers {
		o(val)
	}
}
```

## Walkthrough

The test attaches two closures over `sum`. `SetState(10)` writes `state = 10`,
then calls the first (`sum = 10`) and the second (`sum = 10 + 20 = 30`). Order
is registration order, because `Attach` appends and `range` walks forward.

## Pitfalls

- **Notifying before assigning.** An observer that reads `s.state` would see
  the stale value — a classic source of "the UI is one update behind".
- **Passing `s.state` instead of `val`.** Equivalent here, but it breaks the
  moment a notification is deferred or batched.
- **Notifying in goroutines.** The test's closures share `sum` with no
  synchronization; async delivery would be a data race.
- **Value receiver.** The state write lands on a copy.

## Re-entrancy

If an observer calls `Attach` during notification, it mutates the slice being
ranged. `range` captured the original slice header, so the new observer will not
fire this round — surprising, but at least not a crash. Production buses copy
the slice before iterating.
