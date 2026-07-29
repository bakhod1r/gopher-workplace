# Embedded Field Shadowing

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`Entity` embeds `Base` (which has `ID`) but also declares its own `ID`. The outer
`ID` **shadows** the promoted one, so `e.ID` is the wrong field.

## Task

Fix the return between the markers in
[embeddingshadow.go](embeddingshadow.go) to read the embedded `Base.ID`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Entity{Base:{ID:7}, ID:99}
Output: 7
```

**Example 2:**

```
Input:  Entity{Base:{ID:1}, ID:0}
Output: 1
```

**Example 3:**

```
Input:  Entity{Base:{ID:-3}, ID:5}
Output: -3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Embedding** | Fields promote to the outer type. |
| 2 | **Shadowing** | An outer field hides a promoted one. |
| 3 | **Qualify** | `e.Base.ID` reaches the embedded field. |

## Hint

`return e.Base.ID`.

## Validate

```bash
make verify
```
