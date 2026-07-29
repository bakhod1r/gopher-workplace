# Nil vs Empty Slice

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

A `nil` slice and an empty slice both have length 0, but they encode differently:
`nil` marshals to JSON `null`, an empty slice to `[]`. The function returns nil
when nothing matches, breaking the API contract.

## Task

Fix the declaration between the markers in [nilvsempty.go](nilvsempty.go) to start
from a non-nil empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [a,,b]
Output: [a,b]
```

**Example 2:**

```
Input:  [,,]
Output: [] (non-nil)
```

**Example 3:**

```
Input:  []
Output: [] (non-nil)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nil vs empty** | Same len, different identity. |
| 2 | **JSON encoding** | nil → null, empty → []. |
| 3 | **Init non-nil** | `[]string{}` or `make`. |

## Hint

`out := []string{}`.

## Validate

```bash
make verify
```
