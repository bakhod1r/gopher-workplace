# ARC List Promotion

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Adaptive Replacement Cache keeps two lists: `T1` for items seen once and `T2`
for items seen again. This puzzle models the promotion rule alone — the part
that decides which list an accessed key belongs in.

## Task

Implement `Access` on `*ARC` in [arcpool.go](arcpool.go):

1. If the key is in `T1`, remove it from `T1` and put it in `T2`.
2. If it is already in `T2`, leave it there.
3. Otherwise (unseen) put it in `T1`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Access(1) on an empty ARC
Output: T1 has 1, T2 does not
```

**Example 2:**

```
Input:  Access(1) again
Output: T1 does not have 1, T2 does
```

**Example 3:**

```
Input:  Access(1) a third time
Output: unchanged — still only in T2
```

_Explanation:_ `T2` is terminal for this simplified model.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`delete` on maps** | Promotion is two operations: `delete(a.T1, key)` then `a.T2[key] = true`. |
| 2 | **Set semantics via `map[int]bool`** | The value is irrelevant; membership is the data. |
| 3 | **Branch order** | Check `T2` (or return early after promoting) so a promoted key is not demoted back. |

## Hint

Three branches, and the order matters. `a.T1[key]` on a missing key returns
`false` — the zero value — so a plain lookup is enough for a bool-valued set.

## Validate

```bash
make verify
```
