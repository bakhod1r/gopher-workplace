# Nil Interface Trap

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A request pipeline returns an error interface. A subtle bug makes callers see a failure even when nothing went wrong.

## Task

Implement the stub(s) in [niliface.go](niliface.go):

1. Implement `Run` so it returns a truly nil `error` on success and a `*OpError` on failure.
2. Implement `IsNil`, which reports whether an `error` interface value is nil.
3. Implement `FailedCount`, which counts how many results are non-nil errors.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Run(false)
Output: nil (and IsNil(err) == true)
```

**Example 2:**

```
Input:  Run(true)
Output: error "op failed"
```

**Example 3:**

```
Input:  FailedCount([]error{Run(false), Run(true)})
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Typed nil in an interface** | An interface holding a nil `*OpError` is itself non-nil. |
| 2 | **Interface representation** | An interface value is a (type, value) pair; only both-nil compares equal to nil. |
| 3 | **Error returns** | Reused from error handling: nil means success. |

## Hint

Declare the failure variable inside the failure branch — never return a nil-valued typed pointer as `error`.

## Validate

```bash
make verify
```
