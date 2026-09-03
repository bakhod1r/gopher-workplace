# Wrap Only Real Failures

**Level:** middle
**Topic:** 04-error-handling

## Context

A helper annotates whatever a step returned. Steps that succeeded must keep succeeding.

## Task

Implement `WrapNonNil` in [wrapnonnil.go](wrapnonnil.go):

1. Return nil when `err` is nil, whatever the message says.
2. Return `fmt.Errorf("%s: %w", msg, err)` otherwise.
3. Work when `msg` is empty, producing `": <err>"`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  WrapNonNil("step", nil)
Output: nil
```

**Example 2:**

```
Input:  WrapNonNil("step", ErrX)
Output: "step: boom"
```

**Example 3:**

```
Input:  WrapNonNil("", ErrX)
Output: ": boom"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Typed nil hazard** | A nil error must not become a non-nil wrapper. |
| 2 | **Guard before format** | The check precedes any allocation. |
| 3 | **Honest contracts** | Helpers must not invent failures. |

## Hint

`fmt.Errorf("%s: %w", msg, nil)` returns a perfectly non-nil error — that is the trap.

## Validate

```bash
make verify
```
