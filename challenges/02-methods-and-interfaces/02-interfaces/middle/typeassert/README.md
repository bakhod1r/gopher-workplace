# Type Assertion

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A plugin bus receives opaque payloads and must extract the ones it understands without crashing on the rest.

## Task

Implement the stub(s) in [typeassert.go](typeassert.go):

1. Implement `AsInt`, which returns the int inside the value and whether it was one.
2. Implement `SumInts`, which totals every `int` in the slice and ignores everything else.
3. Never let a bad payload panic — use the comma-ok form.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  AsInt(5)
Output: 5, true
```

**Example 2:**

```
Input:  AsInt("5")
Output: 0, false
```

**Example 3:**

```
Input:  SumInts([]any{1, "x", 2})
Output: 3
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Comma-ok type assertion** | `v, ok := x.(int)` never panics; the single-result form does. |
| 2 | **Zero value on failure** | A failed assertion yields the type's zero value. |
| 3 | **Filtering a heterogeneous slice** | Reused: accumulate only matching elements. |

## Hint

`x.(int)` alone panics on mismatch. `v, ok := x.(int)` does not.

## Validate

```bash
make verify
```
