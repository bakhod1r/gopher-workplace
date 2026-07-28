# Curried Add3

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Currying nests closures: each call captures one argument and returns a function
waiting for the next.

## Task

Implement `Add3` in [curry.go](curry.go).

Do **not** change the function signature or the tests.

## Examples

```go
Add3()(1)(2)(3) // => 6
Add3()(10)(0)(0) // => 10
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nested closures** | Each layer captures one argument. |
| 2 | **Accumulated capture** | Inner functions see all outer arguments. |
| 3 | **Return function chain** | Types nest: func returning func returning func. |

## Hint

`return func(a int) func(int) func(int) int { return func(b int) func(int) int { return func(c int) int { return a+b+c } } }`.

## Validate

```bash
make verify
```
