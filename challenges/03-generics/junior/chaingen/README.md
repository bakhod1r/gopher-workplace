# Fallback Chain

**Level:** junior  
**Topic:** 03-generics

## Context

Configuration is layered: request overrides, then session, then defaults. The first layer holding a key wins.

## Task

Implement the stub(s) in [chaingen.go](chaingen.go):

1. Implement `Lookup`, searching the maps in order and returning the first hit.
2. Return the zero value and `false` when no map holds the key.
3. A stored zero value still counts as a hit.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Lookup(a, {a:1}, {a:2})
Output: 1, true
```

**Example 2:**

```
Input:  Lookup(a, {}, {a:2})
Output: 2, true
```

**Example 3:**

```
Input:  Lookup(a, {}, {})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Variadic maps** | `maps ...map[K]V` accepts any number of layers, all with the same key and value types. |
| 2 | **Comma-ok decides** | Only `ok` distinguishes "layer holds a zero" from "layer lacks the key". |
| 3 | **Priority order** | Returning at the first hit is what makes earlier layers win. |

## Hint

Return at the first `ok` — a stored zero must not fall through to the next layer.

## Validate

```bash
make verify
```
