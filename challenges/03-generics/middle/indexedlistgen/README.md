# List With Fast Membership

**Level:** middle  
**Topic:** 03-generics

## Context

A playlist needs both position-based access and instant "is this track already queued?" checks.

## Task

Implement the stub(s) in [indexedlistgen.go](indexedlistgen.go):

1. Implement `NewIndexed`, `Append`, `Has`, and `At`.
2. `Append` refuses duplicates and reports whether it added the value.
3. `Has` must be constant time — no scanning.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Append(a)
Output: true
```

**Example 2:**

```
Input:  Append(a) twice; Has(a)
Output: true, second Append false
```

**Example 3:**

```
Input:  At(5) on a short list
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Redundant structures** | The slice gives order; the map gives O(1) membership. |
| 2 | **Keeping them in step** | Every mutation must update both, or the invariant breaks. |
| 3 | **Memory cost** | The index roughly doubles the memory — worth it only when lookups are frequent. |

## Hint

Write to the map before appending, so the recorded position matches.

## Validate

```bash
make verify
```
