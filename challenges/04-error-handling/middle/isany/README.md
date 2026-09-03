# Match Any Sentinel

**Level:** middle
**Topic:** 04-error-handling

## Context

A retry policy treats several distinct failures as transient. One helper answers whether an error is in that set.

## Task

Implement `IsAny` in [isany.go](isany.go):

1. Return `true` when `err` matches any of the targets.
2. Return `false` when it matches none of them.
3. Return `false` when there are no targets.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IsAny(ErrTimeout, ErrTimeout, ErrReset)
Output: true
```

**Example 2:**

```
Input:  IsAny(fmt.Errorf("a: %w", ErrReset), ErrTimeout, ErrReset)
Output: true
```

**Example 3:**

```
Input:  IsAny(ErrFatal, ErrTimeout)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Variadic parameters** | `targets ...error` accepts any number of sentinels. |
| 2 | **errors.Is per target** | Each target needs its own chain search. |
| 3 | **Empty variadic** | No targets means nothing can match. |

## Hint

`errors.Is` takes one target at a time — the loop supplies them.

## Validate

```bash
make verify
```
