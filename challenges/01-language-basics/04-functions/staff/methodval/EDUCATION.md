# Method values snapshot the receiver

## The idea

Binding a value-receiver method (`c.Get`) copies the receiver then; the bound function is independent of later changes to `c` — unless you re-bind.

## Why it matters

Method-value binding time matters for callbacks that must reflect a specific object state.

## Watch out

- `f := c.Get` copies `c` now; mutating `c` afterwards doesn't change `f`.
- Re-binding after a mutation captures the newer receiver.
