# fmt.Stringer

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Logs are hard to read when `fmt.Print` outputs raw struct fields like `{1 2}`.
Adding a `String()` method tells `fmt` how to print the struct.

## Task

Implement `String` on `Point` in [fmtstringer.go](fmtstringer.go):

1. Return `"(X,Y)"` (no spaces).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Point{1, 2}.String()
Output: "(1,2)"
```

**Example 2:**

```
Input:  Point{-5, -10}.String()
Output: "(-5,-10)"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **fmt.Stringer interface** | `fmt` packages look for `String() string`. |
| 2 | **Implicit interfaces** | You just write the method; no `implements` keyword. |

## Hint

`fmt.Sprintf("(%d,%d)", p.X, p.Y)`.

## Validate

```bash
make verify
```
