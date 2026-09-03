# Join With Typed Nils

**Level:** staff
**Topic:** 04-error-handling

## Context

A collector receives errors from code that returns typed nil pointers. `errors.Join` keeps them, and the result is a non-nil error reporting nothing.

## Task

Implement `Clean` in [nilsafejoin.go](nilsafejoin.go):

1. Skip entries that are nil or a nil `*OpError`.
2. Join the remaining errors.
3. Return nil when nothing real failed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Clean(nil, (*OpError)(nil))
Output: nil
```

**Example 2:**

```
Input:  Clean(&OpError{"read"})
Output: matches it
```

**Example 3:**

```
Input:  Clean()
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Typed nil in a slice** | Each element is an interface value. |
| 2 | **errors.Join and typed nils** | Only untyped nils are dropped. |
| 3 | **Defensive collection** | Filter at the boundary, once. |

## Hint

`errors.Join` drops entries that are exactly nil — a nil `*OpError` is not one of them.

## Validate

```bash
make verify
```
