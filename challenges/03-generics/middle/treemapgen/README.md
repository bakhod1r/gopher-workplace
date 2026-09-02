# Sorted Map

**Level:** middle  
**Topic:** 03-generics

## Context

A report renders rows in key order and is written to incrementally, so re-sorting after every write would dominate the runtime.

## Task

Implement the stub(s) in [treemapgen.go](treemapgen.go):

1. Implement `NewSorted`, `Set`, `Get`, and `Keys`.
2. Keep the key list sorted as keys are inserted rather than sorting on read.
3. Updating an existing key must not insert it twice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Set(b,1); Set(a,2); Keys()
Output: [a b]
```

**Example 2:**

```
Input:  Set(a,1); Set(a,2); Keys()
Output: [a]
```

**Example 3:**

```
Input:  Get(missing)
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Ordered keys** | `cmp.Ordered` gives both map-key comparability and the `<` needed to place keys. |
| 2 | **Insert into a slice** | Grow by one, shift the tail, then write the new key. |
| 3 | **Sorted on write** | One O(n) insert beats an O(n log n) sort per read. |

## Hint

Insert into the key slice at the right position; only new keys move it.

## Validate

```bash
make verify
```
