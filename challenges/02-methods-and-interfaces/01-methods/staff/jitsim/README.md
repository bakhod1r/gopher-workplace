# JIT Tier-Up

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A tiered runtime interprets code until it proves itself hot, then compiles it.
The policy is a counter per operation and a threshold: cold code stays in the
interpreter, hot code tiers up and stays up.

## Task

Implement `Execute` on `*JIT` in [jitsim.go](jitsim.go):

1. Increment the execution count for `op`.
2. When the count reaches `j.Threshold`, mark the op compiled.
3. Return `Compiled` if the op is compiled (including on the promoting
   execution), otherwise `Interpreted`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  New(3); Execute("loop") four times
Output: interp, interp, jit, jit
```

**Example 2:**

```
Input:  New(2); Execute("a"), Execute("b"), Execute("a")
Output: interp, interp, jit   (counters are per operation)
```

**Example 3:**

```
Input:  New(1); Execute("x")
Output: jit   (the first execution already reaches the threshold)
```

_Explanation:_ the promoting execution is served by the new tier, not the old one.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Per-key counters** | `map[string]int` — a missing key reads as 0, so no initialization branch is needed. |
| 2 | **Sticky state** | Once compiled, always compiled; the count keeps rising but the decision does not flip back. |
| 3 | **Boundary semantics** | `>=` on the threshold is what makes `New(1)` compile on the very first call. |

## Hint

Increment first, then compare. Comparing before the increment shifts every
promotion one execution later and makes `New(1)` never compile.

## Validate

```bash
make verify
```
