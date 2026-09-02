# Set Operations

**Level:** junior  
**Topic:** 03-generics

## Context

A permission check combines the roles a user has directly with the roles their group grants, then compares against what a resource requires.

## Task

Implement the stub(s) in [setopsgen.go](setopsgen.go):

1. Implement `Union`, returning every element of either set.
2. Implement `Intersect`, returning only the elements present in both.
3. Both return a new, non-nil map; neither modifies its inputs.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Union({1,2}, {2,3})
Output: {1,2,3}
```

**Example 2:**

```
Input:  Intersect({1,2}, {2,3})
Output: {2}
```

**Example 3:**

```
Input:  Intersect({1}, {})
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Set-as-map** | Reused from earlier: `map[T]struct{}` is the idiomatic Go set. |
| 2 | **Iterating the smaller side** | For `Intersect`, scanning one set and probing the other is O(len(a)). |
| 3 | **Map keys need `comparable`** | A generic type storing a map must constrain its key parameter. |

## Hint

`Union` adds both; `Intersect` scans one and probes the other.

## Validate

```bash
make verify
```
