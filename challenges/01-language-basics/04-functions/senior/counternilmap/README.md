# Nil Map Write in Helper

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Writing to a nil map panics. A map must be initialised with `make` (or a
literal) before any assignment; reads from nil are fine, writes are not.

## Task

Fix [counternilmap.go](counternilmap.go) so tallying works.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Tally(["a","b","a"])
Output: {a:2, b:1}
```

**Example 2:**

```
Input:  Tally(nil)
Output: {} (empty, non-nil)
```

**Example 3:**

```
Input:  Tally(["x"])
Output: {x:1}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nil vs made map** | `var m map[...]` is nil; writes panic. |
| 2 | **make** | `make(map[string]int)` allocates the map. |
| 3 | **Increment on missing key** | `m[w]++` starts from the zero value once the map exists. |

## Hint

Initialise with `m := make(map[string]int)` (or `map[string]int{}`).

## Validate

```bash
make verify
```
