# Zip With

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Zipping walks two slices in lockstep, stopping at the shorter one, applying a
binary function to each pair.

## Task

Implement `ZipWith` in [zipwith.go](zipwith.go).

Do **not** change the function signature or the tests.

## Examples

```go
ZipWith([]int{1,2,3}, []int{10,20}, add) // => [11 22]
ZipWith(nil, []int{1}, f)                // => []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Lockstep iteration** | Index up to `min(len(a), len(b))`. |
| 2 | **Binary combiner** | `f(a[i], b[i])`. |
| 3 | **Shorter wins** | Extra tail is dropped. |

## Hint

Compute `n := min(len(a), len(b))`; loop `i < n` appending `f(a[i], b[i])`.

## Validate

```bash
make verify
```
