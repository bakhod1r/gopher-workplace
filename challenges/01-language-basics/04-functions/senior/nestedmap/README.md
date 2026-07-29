# Nested Map Inner Init

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

`g[from]` is nil until the inner map is created. Assigning `g[from][to]` writes
into that nil inner map and panics; you must lazily initialise the inner map.

## Task

Fix [nestedmap.go](nestedmap.go) so edges can be added.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Add(g, "a", "b"); g["a"]["b"]
Output: true
```

**Example 2:**

```
Input:  Add on empty graph
Output: no panic
```

**Example 3:**

```
Input:  Add(g, "a", "c") after "a","b"
Output: both present
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nested maps** | The inner map is nil until made. |
| 2 | **Lazy init** | Create `g[from]` on first use. |
| 3 | **comma-ok** | `if _, ok := g[from]; !ok { ... }`. |

## Hint

Init the inner map when absent: `if g[from] == nil { g[from] = map[string]bool{} }`, then `g[from][to] = true`.

## Validate

```bash
make verify
```
