# FizzBuzz Line

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

The order of divisibility checks matters: test the combined case (both 3 and 5)
before the single ones.

## Task

Implement `FizzBuzz` in [fizzbuzz.go](fizzbuzz.go). Use `strconv.Itoa` for the numeric case.

Do **not** change the function signature or the tests.

## Examples

```go
FizzBuzz(3)  // => "Fizz"
FizzBuzz(5)  // => "Buzz"
FizzBuzz(15) // => "FizzBuzz"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **if / else-if ladder** | Check both-divisible first. |
| 2 | **Modulo test** | `n%3 == 0` etc. |
| 3 | **strconv.Itoa** | Convert the fallthrough number to a string. |

## Hint

Check `n%15==0` first, then `n%3`, then `n%5`, else `strconv.Itoa(n)`.

## Validate

```bash
make verify
```
