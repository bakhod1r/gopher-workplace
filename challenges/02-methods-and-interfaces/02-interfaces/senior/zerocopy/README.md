# Zero Copy Slicing

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A parser copied every field out of a large buffer. The copies dominated both allocations and GC pressure.

## Task

Implement the stub(s) in [zerocopy.go](zerocopy.go):

1. Implement `Fields` on `*ZeroCopyParser`, returning sub-slices of the input rather than copies.
2. Implement `CopyFields`, which returns independent copies.
3. Constraint: `Fields` must not allocate per field (only the result slice), and the docs must be honest about the aliasing that creates.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Fields([]byte("a,bb,c"))
Output: three sub-slices aliasing the input
```

**Example 2:**

```
Input:  mutating the input afterwards
Output: the zero-copy fields change too
```

**Example 3:**

```
Input:  CopyFields on the same input
Output: independent copies
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Zero-copy parsing** | Sub-slices share the backing array — no per-field allocation. |
| 2 | **Aliasing as a contract** | The performance win comes with a lifetime obligation for callers. |
| 3 | **Allocation measurement** | Reused: `AllocsPerRun` makes the claim graded. |

## Hint

`data[start:i]` shares memory; `append([]byte(nil), data[start:i]...)` copies.

## Validate

```bash
make verify
```
