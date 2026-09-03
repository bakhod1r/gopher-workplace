# Slicing A Profile By Label

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

`pprof.Do` attaches labels — request kind, tenant, endpoint — to every sample taken while that code runs. Suddenly one profile answers "where does the CPU go *for write requests*", which no amount of function-level aggregation can tell you.

## Task

Implement `SumByLabel` in [profilelabelsum.go](profilelabelsum.go):

1. Group the sample values by the value of `key`.
2. Group samples that lack the key — including those with no labels at all — under `""`.
3. Ignore samples with a non-positive value; return an empty, non-nil map when nothing survives.

## Examples

**Example 1:**

```
Input:  SumByLabel([{kind=read 3} {kind=write 4} {kind=read 2}], "kind")
Output: {read:5 write:4}
```

**Example 2:**

```
Input:  SumByLabel([{kind=read 3} {tenant=acme 4} {nil 1}], "kind")
Output: {read:3 "":5}
```

**Example 3:**

```
Input:  SumByLabel([{kind=read 0}], "kind")
Output: {} (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Profile labels** | `pprof.Do` tags samples with context the call stack does not carry. |
| 2 | **Missing key is a group** | Unlabelled work is a real bucket, not something to drop silently. |
| 3 | **Reading a nil map is safe** | `m[k]` on a nil map returns the zero value; only writing panics. |

## Topics used again

Nested maps, zero values, `range`.

## Hint

`s.Labels[key]` already yields `""` for a missing key *and* for a nil map — no branch needed.

## Validate

```bash
make verify
```
