# Multi Map

**Level:** junior  
**Topic:** 03-generics

## Context

HTTP headers can repeat: one key, several values. A plain map would silently overwrite.

## Task

Implement the stub(s) in [multimapgen.go](multimapgen.go):

1. Implement `NewMultiMap`, `Add`, and `Get`.
2. `Add` keeps every value in insertion order.
3. `Get` returns an empty (non-nil) slice for an unknown key, and a copy for a known one.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Add("a", 1); Add("a", 2); Get("a")
Output: []int{1, 2}
```

**Example 2:**

```
Input:  Get("missing")
Output: []int{}
```

**Example 3:**

```
Input:  Add("a", 1); Add("b", 2); Get("b")
Output: []int{2}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Map of slices** | `map[K][]V` needs both type parameters, each with its own constraint. |
| 2 | **Appending to a missing key** | `append(m.items[k], v)` works because the missing entry reads as a nil slice. |
| 3 | **Defensive copies** | Returning the stored slice would let callers append into the map's storage. |

## Hint

`append(m.items[k], v)` handles the first insert with no special case.

## Validate

```bash
make verify
```
