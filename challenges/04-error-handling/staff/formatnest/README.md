# Verbose Chain

**Level:** staff
**Topic:** 04-error-handling

## Context

With the debug flag on, an error prints its whole chain one line per layer, so a reader sees exactly which layer added what.

## Task

Implement `Verbose` in [formatnest.go](formatnest.go):

1. Return `""` for a nil error.
2. Return one line per error in the chain, outermost first, joined by newlines.
3. Prefix each line after the first with `"caused by: "`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Verbose(ErrA)
Output: "a"
```

**Example 2:**

```
Input:  Verbose(fmt.Errorf("x: %w", ErrA))
Output: "x: a\ncaused by: a"
```

**Example 3:**

```
Input:  Verbose(nil)
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Layer-by-layer output** | Each annotation gets its own line. |
| 2 | **Chain traversal** | Single-error unwrapping only. |
| 3 | **Readable diagnostics** | Structure beats one long line. |

## Hint

The first line has no prefix; every later line does.

## Validate

```bash
make verify
```
