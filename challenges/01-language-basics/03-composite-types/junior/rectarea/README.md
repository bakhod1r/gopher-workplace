# Rectangle Struct

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

A geometry type with methods: area, and a scale that returns a new value without
mutating the receiver.

## Task

Implement `Area()` and `Scale(factor)` on `Rect`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Rect{3,4}.Area()
Output: 12
```

**Example 2:**

```
Input:  Rect{2,3}.Scale(2)
Output: Rect{4,6}
```

**Example 3:**

```
Input:  original after Scale
Output: Rect{2,3}
```

_Explanation:_ Value receiver returns a new Rect; original untouched.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Value receiver** | `func (r Rect)` gets a copy. |
| 2 | **Return new struct** | Scale builds a fresh Rect. |
| 3 | **Struct equality** | `==` compares fields. |

## Hint

`Area`: `r.W * r.H`. `Scale`: `Rect{r.W*factor, r.H*factor}`.

## Validate

```bash
make verify
```
