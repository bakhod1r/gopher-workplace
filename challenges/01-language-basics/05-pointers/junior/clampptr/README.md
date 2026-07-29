# Clamp In Place

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Clamping in place reads and conditionally rewrites the caller's variable through
a pointer.

## Task

Implement `Clamp` in [clampptr.go](clampptr.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  x := 99; Clamp(&x, 0, 10)
Output: x == 10
```

_Explanation:_ Above the range → pulled down to `hi`.

**Example 2:**

```
Input:  x := -5; Clamp(&x, 0, 10)
Output: x == 0
```

_Explanation:_ Below the range → raised to `lo`.

**Example 3:**

```
Input:  x := 4; Clamp(&x, 0, 10)
Output: x == 4
```

_Explanation:_ Inside the range → unchanged.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **In-place clamp** | Rewrite `*p` only when out of range. |
| 2 | **Inclusive bounds** | Endpoints stay. |
| 3 | **Read then write** | `if *p < lo { *p = lo }`. |

## Hint

`if *p < lo { *p = lo } else if *p > hi { *p = hi }`.

## Validate

```bash
make verify
```
