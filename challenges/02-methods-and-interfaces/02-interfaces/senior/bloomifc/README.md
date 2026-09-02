# Bloom Filter

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

Membership checks over a huge key set must fit in a few kilobytes. A false positive is acceptable; a false negative is not.

## Task

Implement the stub(s) in [bloomifc.go](bloomifc.go):

1. Implement `Add` and `MayContain` on `*Bloom` using the two supplied hash functions.
2. Implement `FilterMissing`, which returns the keys the filter says are definitely absent.
3. Constraint: memory is fixed at construction, and the filter must never report a false negative.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Add("a"); MayContain("a")
Output: true
```

**Example 2:**

```
Input:  MayContain on an unadded key
Output: usually false, never a false negative
```

**Example 3:**

```
Input:  100k adds into a fixed bitset
Output: memory unchanged
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Probabilistic data structures** | Trade exactness for a hard memory bound. |
| 2 | **Bitset manipulation** | Index by `bit/8`, mask by `1 << (bit%8)`. |
| 3 | **No false negatives** | The invariant that makes the structure useful as a pre-filter. |

## Hint

Set both hash bits on `Add`; `MayContain` is true only when *both* bits are set.

## Validate

```bash
make verify
```
