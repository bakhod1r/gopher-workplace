# Variadic Max with OK

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _variadic_

## Context

Combining a variadic parameter with an ok flag handles the empty call safely.

## Task

Implement `Max` in [maxvar.go](maxvar.go).

Do **not** change the function signature or the tests.

## Examples

```go
Max()             // => 0, false
Max(3, 9, 2)      // => 9, true
Max(-1, -5)       // => -1, true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Variadic slice** | `nums` is `[]int`. |
| 2 | **Empty guard** | Return 0,false for no args. |
| 3 | **Seed from first** | Start max at `nums[0]`. |

## Hint

Guard `len(nums)==0`; seed `m := nums[0]`; loop the rest updating `m`; return `m, true`.

## Validate

```bash
make verify
```
