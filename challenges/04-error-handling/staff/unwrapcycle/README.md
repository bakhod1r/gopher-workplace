# Survive A Cycle

**Level:** staff
**Topic:** 04-error-handling

## Context

An error type built from user input can be made to unwrap to itself. A traversal that trusts the chain hangs forever.

## Task

Implement `Chain` in [unwrapcycle.go](unwrapcycle.go):

1. Return the messages of every error in the chain, outermost first.
2. Stop when an error repeats, without revisiting it.
3. Return nil for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Chain(fmt.Errorf("x: %w", ErrA))
Output: ["x: a" "a"]
```

**Example 2:**

```
Input:  Chain(selfWrapping)
Output: one entry, terminates
```

**Example 3:**

```
Input:  Chain(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Hostile inputs** | A chain is not guaranteed acyclic. |
| 2 | **Visited sets** | Track identity, not message. |
| 3 | **Comparable keys** | Errors used as map keys must be comparable. |

## Hint

A `map[error]bool` works only for comparable errors — the test's cycle uses a pointer type, which is comparable.

## Validate

```bash
make verify
```
