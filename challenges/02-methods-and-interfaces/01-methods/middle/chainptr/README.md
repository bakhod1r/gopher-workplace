# Chain Pointer

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A configuration builder uses method chaining: `NewConfig().Set("a","1").Set("b","2")`.
This requires each method to return the receiver pointer.

## Task

Implement `Set` on `*Config` in [chainptr.go](chainptr.go):

1. Store `key → value` in `Data`.
2. Return `c` (the pointer) to enable chaining.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NewConfig().Set("host", "localhost").Set("port", "8080")
Output: Data == {"host":"localhost", "port":"8080"}
```

**Example 2:**

```
Input:  NewConfig().Set("a", "1").Set("a", "2")
Output: Data == {"a":"2"} (overwritten)
```

**Example 3:**

```
Input:  NewConfig().Set("x", "y")
Output: Data == {"x":"y"}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method chaining** | Return `*Config` from `Set` to allow `.Set().Set()`. |
| 2 | **Pointer receiver** | Mutation via `c.Data[key] = value`. |
| 3 | **Builder pattern** | Common in Go for configuration and setup. |

## Hint

`c.Data[key] = value; return c`.

## Validate

```bash
make verify
```
