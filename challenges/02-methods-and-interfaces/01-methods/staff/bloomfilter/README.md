# Bloom Filter

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A Bloom filter answers "definitely not present" or "probably present" using
only bits. Each item sets one bit per hash function; a lookup is present only if
every one of its bits is set. False positives are expected; false negatives are
impossible.

## Task

Implement `Add` and `MightContain` on `*Filter` in [bloomfilter.go](bloomfilter.go):

1. `Add` sets `f.bits[hash1(item)]` and `f.bits[hash2(item)]`.
2. `MightContain` returns true only if **both** bits are set.
3. The empty-string guards are already written — leave them.

**Constraint (staff):** no false negatives across 5,000 items, and `Add`/`MightContain` must not allocate.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Add("hello"); MightContain("hello")
Output: true
```

**Example 2:**

```
Input:  MightContain("world")   // 'w' bit unset
Output: false
```

**Example 3:**

```
Input:  MightContain("ho") after Add("hello")
Output: true — a false positive ('h' and 'o' are both set)
```

_Explanation:_ the filter never stores items, only bits, so collisions read as hits.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bits are never cleared** | There is no `Remove`; that is what guarantees no false negatives. |
| 2 | **`&&`, not `\|\|`** | One matching bit is not evidence; all of them must be set. |
| 3 | **`byte` indexes a [256]bool exactly** | The hash range and the array size line up, so no modulo is needed. |

## Hint

Two assignments in `Add`, one `&&` expression in `MightContain`. The `"ho"` case
in the test is a `t.Log`, not a failure — a false positive is correct behaviour.

## Validate

```bash
make verify
```
