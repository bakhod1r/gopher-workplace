# GCD via Euclid Loop

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

Euclid's algorithm repeatedly replaces the pair (a, b) with (b, a mod b) until
b is zero; a is then the GCD. A `for` loop with a swap expresses it cleanly.

## Task

Implement `GCD` in [gcdloop.go](gcdloop.go) using a loop.

Do **not** change the function signature or the tests.

## Examples

```go
GCD(12, 8)  // => 4
GCD(17, 5)  // => 1
GCD(0, 9)   // => 9
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Loop with reassignment** | `a, b = b, a%b` each step. |
| 2 | **Termination** | Stop when `b == 0`. |
| 3 | **Parallel assignment** | Swap without a temp. |

## Hint

`for b != 0 { a, b = b, a%b }; return a`.

## Validate

```bash
make verify
```
