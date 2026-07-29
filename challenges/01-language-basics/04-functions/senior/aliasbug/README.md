# Assignment Shares Backing Array

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _call-by-value_

## Context

`cp := xs` copies the slice HEADER only; both share one backing array, so
`cp[0] = v` also changes `xs[0]`. A real copy needs `make`+`copy` or append.

## Task

Fix [aliasbug.go](aliasbug.go) so the input is not mutated.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  WithFirst([1 2 3], 9)
Output: [9 2 3]; input stays [1 2 3]
```

**Example 2:**

```
Input:  original slice unchanged
Output: true
```

**Example 3:**

```
Input:  WithFirst([5], 0)
Output: [0]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice assignment aliases** | `cp := xs` shares the array. |
| 2 | **Deep copy** | `make`+`copy` or `append([]T(nil), xs...)`. |
| 3 | **Header vs array** | Copying the header is not copying the data. |

## Hint

Make an independent copy: `cp := append([]int(nil), xs...)` (or make+copy).

## Validate

```bash
make verify
```
