# To Set

**Level:** junior  
**Topic:** 03-generics

## Context

A membership check runs inside a hot loop. Converting the allowed list to a set once turns each check from a scan into a lookup.

## Task

Implement the stub(s) in [tosetgen.go](tosetgen.go):

1. Implement `ToSet`, returning a map whose keys are the distinct values of `s`.
2. Return an empty (non-nil) map for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ToSet([]int{1, 1, 2})
Output: keys {1, 2}
```

**Example 2:**

```
Input:  ToSet([]string{"a"})
Output: keys {"a"}
```

**Example 3:**

```
Input:  ToSet([]int{})
Output: empty map
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Maps with type parameters** | A map key type must be comparable, so `K` needs the `comparable` constraint. |
| 2 | **`struct{}` values** | An empty struct occupies zero bytes — the idiomatic set value in Go. |
| 3 | **The `comparable` constraint** | `comparable` is what lets you use `==` and `!=` on a type parameter. |

## Hint

Assigning the same key twice is harmless — no duplicate check needed.

## Validate

```bash
make verify
```
