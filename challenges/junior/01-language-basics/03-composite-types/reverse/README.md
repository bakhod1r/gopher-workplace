# Reverse In Place

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types
**Estimated time:** 10 min

## Context

A rendering step needs a slice flipped end-to-end without allocating a second
slice. Because a slice is a view over a shared backing array, swapping elements
through it changes the caller's data directly.

## Task

Implement `Reverse` in [reverse.go](reverse.go) so it reverses `nums` **in
place** (no new slice, no return value). Empty and single-element slices stay
unchanged.

Do **not** change the function signature or the tests.

## Examples

```go
s := []int{1, 2, 3}; Reverse(s) // s == []int{3, 2, 1}
s := []int{1, 2};    Reverse(s) // s == []int{2, 1}
s := []int{9};       Reverse(s) // s == []int{9}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slices share backing** | A slice passed to a function points at the same array, so element writes are visible to the caller. |
| 2 | **Two-pointer swap** | Swap `nums[i]` and `nums[j]` walking `i` up from the start and `j` down from the end. |
| 3 | **Stop at the middle** | Swap only the first half (`i < len/2`); walking the full length swaps every pair twice and undoes itself. |
| 4 | **Tuple swap** | `nums[i], nums[j] = nums[j], nums[i]` swaps without a temp. |

## Hint

Use two indices, one from each end, swapping and moving inward. Loop while
`i < len(nums)/2` — going the whole length reverses and then re-reverses, so the
slice comes back unchanged.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
