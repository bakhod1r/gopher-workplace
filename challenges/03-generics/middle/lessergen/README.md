# Self-Referential Constraint

**Level:** middle  
**Topic:** 03-generics

## Context

Domain types define their own ordering — version numbers, priorities, semantic ranks — and none of them is a plain number.

## Task

Implement the stub(s) in [lessergen.go](lessergen.go):

1. Implement `MaxOf` using the `Less` method each element provides.
2. On a tie keep the earlier element; return zero and `false` for an empty slice.
3. Study `Lesser[T]`: the constraint mentions the very type it constrains.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MaxOf([]Version{{1}, {3}})
Output: {3}, true
```

**Example 2:**

```
Input:  MaxOf([]Version{{2}, {2}})
Output: the first, true
```

**Example 3:**

```
Input:  MaxOf([]Version{})
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Self-referential constraints** | `[T Lesser[T]]` reads: `T` must be able to compare itself with a `T`. |
| 2 | **Why the parameter is needed** | A plain `Less(any) bool` would lose type safety at the call. |
| 3 | **Replacing `cmp.Ordered`** | This is how ordering is expressed for types that have no `<`. |

## Hint

`best.Less(v)` means "v is bigger" — replace `best` when it holds.

## Validate

```bash
make verify
```
