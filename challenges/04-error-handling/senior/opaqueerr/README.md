# Opaque Error Surface

**Level:** senior
**Topic:** 04-error-handling

## Context

A package exposes only two things about its failures: whether the operation may be retried and whether the input was rejected. Concrete types stay internal.

## Task

Implement `IsRetryable` in [opaqueerr.go](opaqueerr.go):

1. Return `true` from `IsRetryable` for errors carrying `ErrTransient`.
2. Return `true` from `IsRejected` for errors carrying `ErrInvalid`.
3. Return false from both for anything else, including nil.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IsRetryable(fmt.Errorf("x: %w", ErrTransient))
Output: true
```

**Example 2:**

```
Input:  IsRejected(ErrInvalid)
Output: true
```

**Example 3:**

```
Input:  IsRetryable(ErrInvalid)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Opaque errors** | Callers ask questions instead of inspecting types. |
| 2 | **Predicate API** | Behaviour is exported, representation is not. |
| 3 | **Independent predicates** | One error answers exactly one question here. |

## Hint

Two stubs, each one line. The point is the API shape, not the implementation.

## Validate

```bash
make verify
```
