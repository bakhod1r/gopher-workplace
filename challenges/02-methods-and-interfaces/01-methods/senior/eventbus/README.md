# Event Bus

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Publishers should not know their subscribers. A bus keeps a map from event name
to a list of listeners; `On` registers, `Emit` fans out. Everything runs
synchronously, in registration order.

## Task

Implement `On` and `Emit` on `*Bus` in [eventbus.go](eventbus.go):

1. `On(eventType, listener)` appends the listener to `b.listeners[eventType]`.
2. `Emit(eventType, data)` calls every listener registered for that type, in order.
3. Emitting an event with no listeners must not panic.

**Constraint (senior):** a 10,000-listener fan-out must reach every listener, and emitting an event with no listeners must not allocate.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  On("user.login", A); Emit("user.login", "alice")
Output: A receives "alice"
```

**Example 2:**

```
Input:  two listeners on "user.login"; Emit(..., "alice")
Output: both receive "alice", in registration order
```

**Example 3:**

```
Input:  Emit("user.logout", "bob") with nothing registered
Output: no-op, no panic
```

_Explanation:_ a missing map key yields a nil slice, and ranging over nil is legal.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Map of slices** | `append(m[k], v)` works even when `k` is absent — the zero value is a nil slice. |
| 2 | **Ranging a nil slice** | Zero iterations, no panic — that is why `Emit` needs no existence check. |
| 3 | **Functions as values** | Listeners are `func(string)` stored in a slice. |

## Hint

`b.listeners[eventType] = append(b.listeners[eventType], listener)` is the whole
of `On`. `Emit` is a single `range` loop — the "no listeners" case needs no
special handling.

## Validate

```bash
make verify
```
