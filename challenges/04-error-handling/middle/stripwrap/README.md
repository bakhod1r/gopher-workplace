# Report The Known Cause

**Level:** middle
**Topic:** 04-error-handling

## Context

A user-facing layer shows a short, stable message. It replaces the annotated internal error with the sentinel it recognises.

## Task

Implement `Strip` in [stripwrap.go](stripwrap.go):

1. Return `ErrNotFound` or `ErrDenied` when the chain matches one of them.
2. Return `err` unchanged when it matches neither.
3. Return nil for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Strip(fmt.Errorf("a: %w", ErrNotFound))
Output: ErrNotFound
```

**Example 2:**

```
Input:  Strip(errors.New("boom"))
Output: the same error
```

**Example 3:**

```
Input:  Strip(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Boundary translation** | Internal detail stops at the edge. |
| 2 | **errors.Is before replacement** | Only known causes are collapsed. |
| 3 | **Passthrough default** | Unrecognised errors keep their information. |

## Hint

Collapse only what you recognise — anything else must survive untouched, annotations and all.

## Validate

```bash
make verify
```
