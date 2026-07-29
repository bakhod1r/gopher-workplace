# copy Direction for Shift

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

Shifting left drops `xs[0]`: element `i+1` moves to `i`. The code does
`copy(xs[1:], xs)`, which shifts **right** (overlapping copy toward higher
indices), duplicating the first element.

## Task

Fix the copy between the markers in
[copyshiftleft.go](copyshiftleft.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1 2 3]
Output: [2 3 0]
```

**Example 2:**

```
Input:  [5]
Output: [0]
```

**Example 3:**

```
Input:  [] 
Output: []
```

_Explanation:_ empty returns early.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **copy on overlap** | `copy` handles overlap correctly. |
| 2 | **Source/dest** | Left shift: `copy(xs, xs[1:])`. |
| 3 | **Fill vacated** | Zero the last slot. |

## Hint

`copy(xs, xs[1:])`.

## Validate

```bash
make verify
```
