# Method Values

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A callback registered as `obj.Method` kept seeing stale state. The receiver was captured at binding time, not at call time.

## Task

Implement the stub(s) in [methodvalue.go](methodvalue.go):

1. Implement `Get` and `Set` on `*Counter` (pointer receiver) and `Get` on `ValCounter` (value receiver).
2. Implement `BindValue` and `BindPointer`, returning method values bound from a value receiver and a pointer receiver.
3. Constraint: the value-bound closure must keep observing the copy made at bind time, and the pointer-bound one must observe later mutations — the tests pin both.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  bind a value receiver, then mutate
Output: the closure returns the old value
```

**Example 2:**

```
Input:  bind a pointer receiver, then mutate
Output: the closure returns the new value
```

**Example 3:**

```
Input:  a method expression
Output: the receiver is an explicit argument
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method values** | `x.M` copies the receiver *at bind time*; that copy lives in the closure. |
| 2 | **Value versus pointer receivers** | A pointer receiver captures the pointer, so later writes are visible. |
| 3 | **Method expressions** | `T.M` is a function whose first argument is the receiver. |

## Hint

`v.Get` on a value receiver evaluates `v` immediately and stores the copy — the closure never re-reads the variable.

## Validate

```bash
make verify
```
