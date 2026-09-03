# Error With A Line Number

**Level:** middle
**Topic:** 04-error-handling

## Context

A config parser reports where a file went wrong. The line number must reach the caller as data, not buried in prose.

## Task

Implement `LineOf` in [parseerr.go](parseerr.go):

1. Give `*ParseError` an `Error() string` of the form `"line <Line>: <Msg>"`.
2. Implement `LineOf` so it returns the line of a `*ParseError` anywhere in the chain.
3. Return `0, false` when the chain holds no `*ParseError`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  (&ParseError{Line: 4, Msg: "bad token"}).Error()
Output: "line 4: bad token"
```

**Example 2:**

```
Input:  LineOf(&ParseError{Line: 4})
Output: 4, true
```

**Example 3:**

```
Input:  LineOf(errors.New("boom"))
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Structured errors** | Fields survive where formatted text does not. |
| 2 | **errors.As** | Recovers the typed error through wrapping. |
| 3 | **Pointer receivers** | `*ParseError` is the error implementation. |

## Hint

Two stubs: the message format, and the `errors.As` lookup that pulls the struct back out.

## Validate

```bash
make verify
```
