# Verifying Stability

**Level:** middle  
**Topic:** 03-generics

## Context

A regression test must prove a sort is stable, not merely ordered, because the UI depends on equal rows keeping their positions.

## Task

Implement the stub(s) in [sortstablecheckgen.go](sortstablecheckgen.go):

1. Implement `IsStableBy`, checking both the ordering and the relative order within equal keys.
2. Compare against the original slice to know what the input order was.
3. Assume the elements are distinct enough to be used as map keys.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  stable sort output
Output: true
```

**Example 2:**

```
Input:  reversed equal group
Output: false
```

**Example 3:**

```
Input:  empty slice
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Stability is not order** | A sorted result can still be unstable — the extra check is what proves it. |
| 2 | **Recording original positions** | A map from element to first index gives the input order. |
| 3 | **Two separate checks** | Ordering and stability fail for different reasons; check them separately. |

## Hint

Sortedness alone is not stability — compare equal-key neighbours against their original positions.

## Validate

```bash
make verify
```
