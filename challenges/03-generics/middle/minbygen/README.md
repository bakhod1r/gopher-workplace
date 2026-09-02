# Min By Key

**Level:** middle  
**Topic:** 03-generics

## Context

A dispatcher picks the least-loaded worker. Workers are structs with no natural ordering.

## Task

Implement the stub(s) in [minbygen.go](minbygen.go):

1. Implement `MinBy`, returning the element with the smallest key and `true`.
2. On a tie keep the earlier element; return the zero value and `false` for an empty slice.
3. Call `key` at most once per element.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MinBy(workers, loadOf)
Output: least loaded worker, true
```

**Example 2:**

```
Input:  MinBy([]int{3,1,1}, self)
Output: 1 (the first one), true
```

**Example 3:**

```
Input:  MinBy([]int{}, self)
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Key projections** | A `func(T) K` decouples what to compare from how to traverse. |
| 2 | **Caching the current best key** | Storing `bestKey` keeps the number of `key` calls linear. |
| 3 | **`cmp.Ordered`** | The stdlib constraint for `<`, `<=`, `>`, `>=`. |

## Hint

Seed from `s[0]` and cache its key.

## Validate

```bash
make verify
```
