# Binding method values on pointers vs values

## The idea

`a.Get` (pointer) binds the live object; `(*a).Get` binds a copy of the value, freezing its state at bind time.

## Why it matters

Accidentally binding on a dereferenced value produces stale getters — a subtle bug.

## Watch out

- `(*a).Get` snapshots the value; `a.Get` tracks the pointer.
- Prefer binding the method value on the pointer for live views.
