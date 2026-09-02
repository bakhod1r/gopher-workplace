# Once With Result

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An initialiser that panicked left the cache in a half-built state, and every later caller saw a different failure. Initialisation must be exactly-once, including its failure.

## Task

Implement the stub(s) in [onceresult.go](onceresult.go):

1. Implement `Get` on `*OnceValue`, running the initialiser once and caching its value and error.
2. A panic inside the initialiser must be converted into a cached error, not left to escape on every call.
3. Constraint: `-race` clean; under concurrent first calls the initialiser runs exactly once, and every caller observes the same result.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  two Get calls
Output: the initialiser runs once
```

**Example 2:**

```
Input:  an initialiser that fails
Output: the same error every time, one run
```

**Example 3:**

```
Input:  an initialiser that panics
Output: a cached error, no panic escapes
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Once with results** | `Once` returns nothing, so the result must be stored in the closure's scope. |
| 2 | **Panic containment** | `recover` inside the once function turns a panic into a value. |
| 3 | **Cached failure** | Reused: retrying a failed init is a policy decision, not a default. |

## Hint

`defer func() { if r := recover(); r != nil { err = ... } }()` inside the once function.

## Validate

```bash
make verify
```
