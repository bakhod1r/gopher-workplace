# Constant Errors

**Level:** staff
**Topic:** 04-error-handling

## Context

A package wants sentinels that cannot be reassigned by a caller, so they are declared as constants rather than variables.

## Task

Implement `Error` in [consterr.go](consterr.go):

1. Give the `Error` string type an `Error() string` method returning itself.
2. Make `ErrClosed` and `ErrBusy` usable as `error` values.
3. Keep them comparable and matchable with `errors.Is`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ErrClosed.Error()
Output: "closed"
```

**Example 2:**

```
Input:  errors.Is(fmt.Errorf("x: %w", ErrClosed), ErrClosed)
Output: true
```

**Example 3:**

```
Input:  ErrClosed == ErrBusy
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Named string types** | A string type can implement `error`. |
| 2 | **Immutable sentinels** | Constants cannot be reassigned by callers. |
| 3 | **Value comparison** | String-typed errors compare by content. |

## Hint

The method is one line; the interesting part is that a constant can satisfy an interface at all.

## Validate

```bash
make verify
```
