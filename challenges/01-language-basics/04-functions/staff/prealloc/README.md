# Prealloc Length vs Append

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

`make([]int, n)` creates a slice of LENGTH n (n zeros); appending then adds
AFTER those zeros. To preallocate capacity without content, use
`make([]int, 0, n)`.

## Task

Fix [prealloc.go](prealloc.go) so the result has no leading zeros.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Doubles(3)
Output: [2 4 6]
```

**Example 2:**

```
Input:  Doubles(0)
Output: []
```

**Example 3:**

```
Input:  Doubles(1)
Output: [2]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **length vs capacity in make** | `make([]T, n)` has n elements already. |
| 2 | **Append after prealloc** | It adds beyond the initial length. |
| 3 | **Zero-length, n-capacity** | `make([]T, 0, n)` reserves without filling. |

## Hint

Preallocate capacity only: `out := make([]int, 0, n)` (then append), or keep length n and assign `out[i-1] = i*2`.

## Validate

```bash
make verify
```
