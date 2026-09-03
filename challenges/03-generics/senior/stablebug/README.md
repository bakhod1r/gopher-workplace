# Rows That Jump Between Renders

**Level:** senior  
**Topic:** 03-generics

## Context

A table sorted by age reshuffles rows of the same age on every refresh, and the snapshot tests are flaky.

## Task

Fix the single planted bug in [stablebug.go](stablebug.go):

1. Find and fix the single bug so equal keys keep their input order.
2. Leave the input untouched.
3. Do not change the comparison itself.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  sort [{a 20} {b 20} {c 10}] by age
Output: [{c 10} {a 20} {b 20}]
```

**Example 2:**

```
Input:  input after the call
Output: unchanged
```

**Example 3:**

```
Input:  SortedBy(nil, ageOf)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Stable versus unstable** | `slices.SortFunc` may reorder equal elements; `SortStableFunc` may not. |
| 2 | **Same comparison, different guarantee** | The bug is in the chosen sort, not in the comparator. |
| 3 | **Small inputs hide it** | Unstable sorts often happen to be stable for short slices. |

## Hint

Both sorts use the same comparison. What else differs?

## Validate

```bash
make verify
```
