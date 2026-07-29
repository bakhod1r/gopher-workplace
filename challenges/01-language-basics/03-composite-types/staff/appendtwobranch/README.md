# Two Appends Share Capacity

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

`append(base, x)` and `append(base, y)` both write into `base`'s spare capacity at
the same index, so the second overwrites the first — `b` ends up `[1 2 4]`.

## Task

Fix the base between the markers in
[appendtwobranch.go](appendtwobranch.go) so the two appends don't alias.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a=[1 2] (cap 10), x=3, y=4
Output: b=[1 2 3], c=[1 2 4]
```

_Explanation:_ each branch must not reuse a's spare capacity, or the second append overwrites the first.

**Example 2:**

```
Input:  b after Branch
Output: [1 2 3]
```

**Example 3:**

```
Input:  c after Branch
Output: [1 2 4]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shared spare cap** | Both appends target the same slot. |
| 2 | **Clip capacity** | `a[:len(a):len(a)]` forces realloc. |
| 3 | **Independent results** | Each append gets its own array. |

## Hint

`base := a[:len(a):len(a)]`.

## Validate

```bash
make verify
```
