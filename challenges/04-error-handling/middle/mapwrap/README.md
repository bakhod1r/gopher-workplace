# Contextual Lookup Failure

**Level:** middle
**Topic:** 04-error-handling

## Context

A repository reports which identifier was missing, while callers keep matching on one shared sentinel.

## Task

Implement `Find` in [mapwrap.go](mapwrap.go):

1. Return the stored value and nil when the key is present.
2. Return `0` and an error reading `"key <key>: not found"` when it is absent.
3. Keep `ErrNotFound` matchable through the wrapper.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Find(map[string]int{"a": 1}, "a")
Output: 1, nil
```

**Example 2:**

```
Input:  Find(nil, "a")
Output: 0, "key a: not found"
```

**Example 3:**

```
Input:  errors.Is(err, ErrNotFound)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sentinel plus context** | One matchable cause, many specific messages. |
| 2 | **Comma-ok idiom** | Presence is still decided by `ok`. |
| 3 | **Wrapping at the source** | Context is cheapest where the failure happens. |

## Hint

The message must name the key, and `errors.Is` must still find the sentinel — one `fmt.Errorf` does both.

## Validate

```bash
make verify
```
