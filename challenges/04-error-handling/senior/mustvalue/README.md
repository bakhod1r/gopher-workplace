# Must Helper

**Level:** senior
**Topic:** 04-error-handling

## Context

Package initialisation loads assets that are supposed to be present. A missing one is a programming error, not a runtime condition.

## Task

Implement `Must` in [mustvalue.go](mustvalue.go):

1. Return `v` when `err` is nil.
2. Panic with the error itself when `err` is non-nil.
3. Panic with the error value, not a string.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Must(42, nil)
Output: 42
```

**Example 2:**

```
Input:  Must(0, ErrLoad)
Output: panics with ErrLoad
```

**Example 3:**

```
Input:  recover() == ErrLoad
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Must convention** | Reserved for failures that cannot be handled. |
| 2 | **Panic payloads** | Any value can be panicked, including an error. |
| 3 | **Initialisation-time failure** | Crashing early beats running broken. |

## Hint

The test recovers and compares the panic value against the sentinel — panic with the error, not `err.Error()`.

## Validate

```bash
make verify
```
