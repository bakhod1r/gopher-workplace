# Copy at the Right Offset

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

The second `copy` writes `b` at offset 0, overwriting `a`. It must write at offset
`len(a)` so `b` lands after `a`.

## Task

Fix the second copy between the markers in [bufconcat.go](bufconcat.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a=[1 2], b=[3 4]
Output: [1 2 3 4]
```

**Example 2:**

```
Input:  a=[], b=[9]
Output: [9]
```

**Example 3:**

```
Input:  a=[5], b=[]
Output: [5]
```

_Explanation:_ copy target offset must be len(a).

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Copy destination** | `out[len(a):]` targets the tail. |
| 2 | **Preallocated buffer** | Size = len(a)+len(b). |
| 3 | **Offset math** | b starts where a ends. |

## Hint

`copy(out[len(a):], b)`.

## Validate

```bash
make verify
```
