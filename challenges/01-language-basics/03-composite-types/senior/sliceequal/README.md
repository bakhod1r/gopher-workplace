# Slice Equality

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

Slices can't be compared with `==` (that's a compile error), so equality must be
element-wise. The code checks only the length, so different contents of equal
length compare "equal".

## Task

Fix the body between the markers in [sliceequal.go](sliceequal.go) to compare
elements.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a=[1,2,3], b=[1,2,3]
Output: true
```

**Example 2:**

```
Input:  a=[1,2,3], b=[1,9,3]
Output: false
```

**Example 3:**

```
Input:  a=[], b=[]
Output: true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **No slice ==** | Only nil comparison is allowed. |
| 2 | **Element-wise** | Loop and compare each index. |
| 3 | **Length first** | Fast reject on differing lengths. |

## Hint

`for i := range a { if a[i] != b[i] { return false } }; return true`.

## Validate

```bash
make verify
```
