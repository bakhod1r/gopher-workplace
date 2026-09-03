# Comparison That Cannot Panic

**Level:** staff
**Topic:** 04-error-handling

## Context

A framework compares caller-supplied errors. An uncomparable concrete type makes the standard comparison panic, taking the request with it.

## Task

Implement `SafeIs` in [safeis.go](safeis.go):

1. Return the result of `errors.Is(err, target)` when it completes.
2. Return false instead of panicking when the comparison panics.
3. Return false when either argument is nil.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SafeIs(ErrA, ErrA)
Output: true
```

**Example 2:**

```
Input:  SafeIs(uncomparable, uncomparable)
Output: false, no panic
```

**Example 3:**

```
Input:  SafeIs(nil, ErrA)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Comparability** | Comparing structs with slice fields panics at runtime. |
| 2 | **Defensive boundaries** | Framework code cannot trust caller types. |
| 3 | **recover around a call** | The guard is local to one comparison. |

## Hint

A struct error with a slice field is uncomparable — `errors.Is` panics when it reaches the `==`.

## Validate

```bash
make verify
```
