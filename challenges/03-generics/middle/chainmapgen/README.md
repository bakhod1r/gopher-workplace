# Layered Configuration

**Level:** middle  
**Topic:** 03-generics

## Context

Settings come from flags, then environment, then a file. Both a live lookup and a flattened snapshot are needed.

## Task

Implement the stub(s) in [chainmapgen.go](chainmapgen.go):

1. Implement `NewChain`, `Get`, and `Flatten`.
2. Earlier layers take priority.
3. `Flatten` returns a new map and never modifies a layer.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Get from {a:1} over {a:2}
Output: 1, true
```

**Example 2:**

```
Input:  Get(missing)
Output: zero, false
```

**Example 3:**

```
Input:  Flatten of {a:1} over {a:2, b:3}
Output: {a:1, b:3}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Priority order** | `Get` searches forward; `Flatten` must therefore merge backward. |
| 2 | **A stored zero is a hit** | Comma-ok is what makes a deliberately zero setting shadow lower layers. |
| 3 | **Variadic layers** | Any number of maps, all with the same key and value types. |

## Hint

`Flatten` merges from the last layer to the first, so the highest-priority values are written last.

## Validate

```bash
make verify
```
