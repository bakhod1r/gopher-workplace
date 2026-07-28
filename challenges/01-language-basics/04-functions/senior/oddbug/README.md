# Parity Test for Negatives

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _conditionals_

## Context

Go's `%` keeps the sign of the dividend, so `-3 % 2 == -1`, not 1. Testing
`== 1` misses negative odds; `!= 0` is the sign-safe parity test.

## Task

Fix [oddbug.go](oddbug.go) so negative odd numbers are counted.

Do **not** change the function signature or the tests.

## Examples

```go
CountOdd([1 2 3 -3 -4 -5]) // => 4
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Signed modulo** | `-3 % 2 == -1` in Go. |
| 2 | **Sign-safe parity** | `v%2 != 0` catches all odds. |
| 3 | **Boundary of the spec** | Negatives must be handled. |

## Hint

Use `v%2 != 0` (equivalently `v%2 == 1 || v%2 == -1`).

## Validate

```bash
make verify
```
