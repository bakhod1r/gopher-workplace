# Re-panic Selectively

**Level:** staff
**Topic:** 04-error-handling

## Context

A request handler recovers to log the failure, but genuine runtime corruption must keep unwinding rather than be turned into a 500.

## Task

Implement `Handle` in [repanic.go](repanic.go):

1. Return an error wrapping the recovered value when it is not a `runtime.Error`.
2. Re-panic with the original value when it is a `runtime.Error`.
3. Return nil when `f` does not panic.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Handle(func() { panic("app bug") })
Output: an error
```

**Example 2:**

```
Input:  Handle(func() { _ = []int{}[3] })
Output: re-panics
```

**Example 3:**

```
Input:  Handle(func() {})
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **runtime.Error** | Distinguishes runtime faults from application panics. |
| 2 | **Re-panicking** | Passing the original value preserves the payload. |
| 3 | **Recovery policy** | Not every panic should be absorbed. |

## Hint

Re-panic with the recovered value itself, not a new one, so the payload the caller sees is unchanged.

## Validate

```bash
make verify
```
