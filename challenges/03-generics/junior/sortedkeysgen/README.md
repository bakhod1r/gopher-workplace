# Sorted Keys

**Level:** junior  
**Topic:** 03-generics

## Context

A config dump must be reproducible across runs, so its keys cannot come out in Go's randomised map order.

## Task

Implement the stub(s) in [sortedkeysgen.go](sortedkeysgen.go):

1. Implement `SortedKeys`, returning the map's keys in ascending order.
2. Return an empty (non-nil) slice for an empty map.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SortedKeys(map[string]int{"b": 1, "a": 2})
Output: []string{"a", "b"}
```

**Example 2:**

```
Input:  SortedKeys(map[int]bool{3: true, 1: false})
Output: []int{1, 3}
```

**Example 3:**

```
Input:  SortedKeys(map[string]int{})
Output: []string{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Ordered` implies `comparable`** | Every ordered type is comparable, so `K` is still a legal map key type. |
| 2 | **Deterministic output** | Reused from earlier: map range order is randomised, so sorting is the fix. |
| 3 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<`, `<=`, `>`, `>=`. |

## Hint

Collect first, sort second — a map cannot be ranged in order.

## Validate

```bash
make verify
```
