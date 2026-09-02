# Command Pattern

## Intuition

Normally calling a function *is* running it. The command pattern splits those:
`Add` stores the call, `ExecuteAll` performs it. Between the two, the work is an
ordinary value you can count, reorder, log or drop.

In Go the pattern needs no interface — a `func()` is already a first-class
value, so `type Command func()` is the whole abstraction.

## Approach

1. Iterate the slice in order and call each element.
2. Reset the slice so the commands cannot run twice.

## Solution

```go
func (inv *Invoker) ExecuteAll() {
	for _, c := range inv.commands {
		c()
	}
	inv.commands = nil
}
```

## Walkthrough

The two closures in the test both capture `x` by reference — they close over the
variable, not its value. Running them in queue order gives `x = 0 + 5 = 5`, then
`x = 5 * 2 = 10`.

`inv.commands = nil` sets the slice header to zero. `len(nil slice) == 0`, so
the test's queue-cleared check passes, and a later `Add` still works because
`append` allocates on a nil slice.

## Pitfalls

- **Clearing first.** `inv.commands = nil` before the loop makes `ExecuteAll` a
  no-op — the loop then ranges over nothing.
- **Value receiver.** The commands run, but the caller's queue is never cleared.
- **`inv.commands = inv.commands[:0]`.** Also gives length 0, but keeps the
  backing array (and its captured closures) alive; `nil` releases them.
