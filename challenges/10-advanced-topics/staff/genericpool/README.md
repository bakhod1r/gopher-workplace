# A Typed Pool With No Assertions

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

Every pool in the codebase is a `sync.Pool` plus a type assertion at each call site. One of them asserts the wrong type, and the panic only fires under the load that fills the pool.

## Task

Implement [genericpool.go](genericpool.go):

1. Return a pointer to a zeroed `T` from the pool.
2. A recycled value must come back zeroed — callers must never see the previous holder's data.
3. No type assertion may escape to the caller; safe for concurrent use.

Replace the stub body in [genericpool.go](genericpool.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  NewPool[buffer]().Get()
Output: a zeroed *buffer
```

**Example 2:**

```
Input:  Put a dirty value, then Get
Output: zeroed again
```

**Example 3:**

```
Input:  NewPool[int]()
Output: works with any type
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Generics over sync.Pool** | The assertion happens once, inside the wrapper. |
| 2 | **Pooled values carry state** | Resetting is the wrapper's job, not the caller's. |
| 3 | **The zero value as the reset** | `*v = zero` works for any T without knowing its fields. |
| 4 | **sync.Pool may drop entries** | `New` covers the empty case, so `Get` never returns nil. |

## Hint

The assertion is safe because `New` is the only thing that ever puts a value in — and then one more line makes it clean.

## Validate

```bash
make verify
```
