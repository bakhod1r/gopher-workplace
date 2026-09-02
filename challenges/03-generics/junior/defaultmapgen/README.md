# Default Map

**Level:** junior  
**Topic:** 03-generics

## Context

A feature-flag store answers every question, returning the configured default for flags nobody has set.

## Task

Implement the stub(s) in [defaultmapgen.go](defaultmapgen.go):

1. Implement `NewDefaultMap`, `Put`, and `Get`.
2. `Get` returns the stored value when present — including a stored zero value — and the fallback otherwise.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NewDefaultMap[string](7); Get("x")
Output: 7
```

**Example 2:**

```
Input:  Put("x", 1); Get("x")
Output: 1
```

**Example 3:**

```
Input:  Put("x", 0); Get("x")
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **State inside a generic type** | The fallback is stored once in the struct instead of passed on every call. |
| 2 | **Comma-ok again** | Reused from earlier: only `ok` can distinguish a stored zero from a missing key. |
| 3 | **Map keys need `comparable`** | A generic type storing a map must constrain its key parameter. |

## Hint

Same comma-ok rule as the `GetOr` function — the default now lives in the struct.

## Validate

```bash
make verify
```
