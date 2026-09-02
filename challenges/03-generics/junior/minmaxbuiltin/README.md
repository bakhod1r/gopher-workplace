# The min And max Builtins

**Level:** junior  
**Topic:** 03-generics

## Context

A layout engine clamps and measures constantly. Go's builtins remove a helper package from the codebase.

## Task

Implement the stub(s) in [minmaxbuiltin.go](minmaxbuiltin.go):

1. Implement `Middle` with the `min` and `max` builtins.
2. Implement `Spread`, returning the largest minus the smallest of three values.
3. Do not write your own comparison helpers here.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Middle(5, 0, 3)
Output: 3
```

**Example 2:**

```
Input:  Middle(-1, 0, 3)
Output: 0
```

**Example 3:**

```
Input:  Spread(1, 9, 4)
Output: 8
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`min` and `max` builtins** | Built into the language since Go 1.21 — variadic, ordered types only. |
| 2 | **Builtins versus `slices.Max`** | The builtins take arguments; `slices.Max` takes a slice. |
| 3 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<` and `>`. |

## Hint

`min(max(v, lo), hi)` clamps in one expression.

## Validate

```bash
make verify
```
