# Binding method values on pointers vs values

## Intuition

`a.Get` (pointer) binds the live object; `(*a).Get` binds a copy of the value, freezing its state at bind time.

## Approach

1. `tmp := *a` copies the struct, so `tmp.Get` is bound to a stale snapshot.
2. Bind the method value to the live pointer: `return a.Get`.

## Solution

```go
type Account struct{ Balance int }

func (a *Account) Get() int { return a.Balance }

func Getter(a *Account) func() int {
	return a.Get
}
```

## Walkthrough

The bug's `tmp.Get` captures a copy taken at bind time; later writes to `*a` are invisible. Binding `a.Get` keeps the pointer, so `get()` sees `250`.

## Pitfalls

- `(*a).Get` snapshots the value; `a.Get` tracks the pointer.
- Prefer binding the method value on the pointer for live views.
