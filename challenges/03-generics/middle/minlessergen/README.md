# Minimum By Less

**Level:** middle  
**Topic:** 03-generics

## Context

The same release list also needs its oldest entry, and only the type itself knows what "older" means.

## Task

Implement the stub(s) in [minlessergen.go](minlessergen.go):

1. Implement `MinOf` using the elements' `Less` method.
2. On a tie keep the earlier element; return zero and `false` for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MinOf([]Version{{3}, {1}})
Output: {1}, true
```

**Example 2:**

```
Input:  MinOf([]Version{{2}, {2}})
Output: the first, true
```

**Example 3:**

```
Input:  MinOf([]Version{})
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Receiver order matters** | `v.Less(best)` and `best.Less(v)` are the two directions — pick deliberately. |
| 2 | **Strictness keeps ties stable** | Replacing only on a strict `Less` preserves the earlier element. |
| 3 | **Self-referential constraint** | Same `Lesser[T]` shape as the maximum. |

## Hint

Flip the receiver and the argument compared with `MaxOf`.

## Validate

```bash
make verify
```
