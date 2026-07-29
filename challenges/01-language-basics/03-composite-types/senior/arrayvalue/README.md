# Arrays Are Values

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`b := *a` copies the whole array (arrays are values). Doubling `b` mutates the
copy; the caller's array via `a` is never touched.

## Task

Fix the body between the markers in [arrayvalue.go](arrayvalue.go) to mutate
through the pointer.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,2,3]
Output: [2,4,6]
```

**Example 2:**

```
Input:  [0,0,0]
Output: [0,0,0]
```

**Example 3:**

```
Input:  [-1,5,10]
Output: [-2,10,20]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Array vs slice** | Arrays copy on assignment; slices share. |
| 2 | **Deref to mutate** | Work through `a`, not a copy of `*a`. |
| 3 | **Auto-index** | `a[i]` indexes through `*[3]int`. |

## Hint

`for i := range a { a[i] *= 2 }`.

## Validate

```bash
make verify
```
