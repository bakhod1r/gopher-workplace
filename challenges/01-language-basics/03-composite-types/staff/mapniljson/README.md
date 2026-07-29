# Nil Map JSON

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

A nil map marshals to JSON `null`; an empty map to `{}`. `Counts` starts with a
nil map and only allocates inside the loop, so **empty input** returns nil → `null`,
breaking clients expecting an object.

## Task

Fix the declaration between the markers in [mapniljson.go](mapniljson.go) to
allocate upfront.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  xs=[] (empty)
Output: non-nil empty map -> JSON {}
```

**Example 2:**

```
Input:  xs=[a a b]
Output: {a:2, b:1}
```

**Example 3:**

```
Input:  xs=[x]
Output: {x:1}
```

_Explanation:_ empty input must not encode to null.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nil vs empty map** | Both len 0; encode differently. |
| 2 | **JSON encoding** | nil → null, empty → {}. |
| 3 | **Allocate upfront** | `make(map[string]int)`. |

## Hint

`m := make(map[string]int)`.

## Validate

```bash
make verify
```
