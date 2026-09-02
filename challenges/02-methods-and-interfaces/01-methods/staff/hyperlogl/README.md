# HyperLogLog Register

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

HyperLogLog estimates how many *distinct* items a stream contained, in constant
memory, by remembering only the longest run of leading zeros ever seen in a
hash. A long run is rare, so seeing one implies many distinct items passed by.

## Task

Implement `Add` on `*HLL` in [hyperlogl.go](hyperlogl.go):

1. Compute `leadingZeros(hash)`.
2. If that is greater than `h.maxZeros`, store it.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Add(1)  // leadingZeros = 1
Output: maxZeros == 1
```

**Example 2:**

```
Input:  Add(4)  // leadingZeros = 4
Output: maxZeros == 4
```

**Example 3:**

```
Input:  Add(2)  // leadingZeros = 2, not greater
Output: maxZeros stays 4
```

_Explanation:_ the register only ever moves up — the estimate must not depend on arrival order.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Monotone register** | `>` and never `=` keeps `Add` idempotent and order-independent. |
| 2 | **Constant memory** | One int summarises an unbounded stream. |
| 3 | **Pointer receiver** | The register must survive the call. |

## Hint

The mock `leadingZeros` is `v % 5` — deliberately fake, so the test values are
easy to reason about. The comparison is the part being tested.

## Validate

```bash
make verify
```
