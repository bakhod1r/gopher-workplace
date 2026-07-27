# Array Value Semantics

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types
**Estimated time:** 10 min

## Context

Unlike a slice, a Go **array** is a value. Assigning it or passing it to a
function copies the whole thing, so a helper can't change the caller's array in
place — it has to return the new one.

## Task

Implement `SetFirst` in [arrayval.go](arrayval.go) so it returns a copy of `a`
with the first element set to `v`, leaving the caller's original array
unchanged.

Do **not** change the function signature or the tests.

## Examples

```go
SetFirst([3]int{1, 2, 3}, 9) // => [3]int{9, 2, 3}
SetFirst([3]int{0, 0, 0}, 5) // => [3]int{5, 0, 0}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Arrays are values** | `[3]int` has a fixed length that is part of its type; copying it copies every element. |
| 2 | **Pass-by-copy** | A function receives its own copy of an array argument; writing to it never touches the caller's. |
| 3 | **Array vs slice** | A slice would share a backing array (writes leak to the caller); an array does not. |

## Hint

Inside the function `a` is already your own copy. Set `a[0] = v` and return
`a`; the caller's array is a separate value and stays as it was.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
