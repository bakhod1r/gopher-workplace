# Even/Odd Label

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _conditionals_

## Context

Testing evenness for negative numbers needs `== 0`, not `== 1`, since `-3 % 2`
is `-1` in Go.

## Task

Implement `Parity` in [gradeif.go](gradeif.go).

Do **not** change the function signature or the tests.

## Examples

```go
Parity(4)  // => "even"
Parity(7)  // => "odd"
Parity(-3) // => "odd"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **if / else** | Two-way branch on the remainder. |
| 2 | **Signed modulo** | `n%2 == 0` is the safe evenness test. |
| 3 | **Return strings** | One of two labels. |

## Hint

`if n%2 == 0 { return "even" }; return "odd"`.

## Validate

```bash
make verify
```
