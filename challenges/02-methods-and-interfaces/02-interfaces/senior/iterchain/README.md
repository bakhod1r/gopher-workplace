# Iterator Chain

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A query layer chains map and filter steps over a large source. Building an intermediate slice per step blew the memory budget.

## Task

Implement the stub(s) in [iterchain.go](iterchain.go):

1. Implement `Next` on `*MapIter` and `*FilterIter`, each pulling from the wrapped iterator lazily.
2. Implement `Collect`, which drains an iterator into a slice.
3. Constraint: chaining must be lazy — no step may materialise its input, and the test asserts the source is read at most once per element.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  map(double) over [1 2 3], collected
Output: [2 4 6]
```

**Example 2:**

```
Input:  filter(even) after map(double)
Output: all elements pass
```

**Example 3:**

```
Input:  a chain that is never collected
Output: the source is never read
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Lazy iterators** | Pull-based chaining keeps memory at O(1) per stage. |
| 2 | **Composable interfaces** | Every stage is both an `Iter` and a consumer of one. |
| 3 | **Deferred evaluation** | Building the chain does no work until it is drained. |

## Hint

`FilterIter.Next` loops until the wrapped iterator yields a matching element or drains.

## Validate

```bash
make verify
```
