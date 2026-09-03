# Errors By Field

**Level:** middle
**Topic:** 04-error-handling

## Context

A request validator returns a map from field name to failure. The response needs one combined error naming each field.

## Task

Implement `Combine` in [fielderrs.go](fielderrs.go):

1. Wrap each non-nil entry as `"<field>: <err>"` and combine them.
2. Process fields in sorted key order so the message is deterministic.
3. Return nil for an empty or all-nil map.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Combine(map[string]error{"b": ErrB, "a": ErrA})
Output: "a: bad a\nb: bad b"
```

**Example 2:**

```
Input:  Combine(nil)
Output: nil
```

**Example 3:**

```
Input:  errors.Is(combined, ErrA)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Map iteration order** | Ranging a map is randomised; sort the keys. |
| 2 | **errors.Join formatting** | Joined messages are separated by newlines. |
| 3 | **Wrap then join** | Context per field, one error overall. |

## Hint

Map range order is deliberately random in Go — collect the keys and sort them before building anything.

## Validate

```bash
make verify
```
