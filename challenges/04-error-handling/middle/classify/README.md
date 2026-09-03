# Failure To Status

**Level:** middle
**Topic:** 04-error-handling

## Context

An API gateway turns internal failures into status codes. Every unknown failure must land on 500 rather than leaking upward untranslated.

## Task

Implement `Status` in [classify.go](classify.go):

1. Return `404` when the chain contains `ErrNotFound`.
2. Return `403` when it contains `ErrDenied`, and `409` for `ErrConflict`.
3. Return `500` for anything else, and `200` for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Status(nil)
Output: 200
```

**Example 2:**

```
Input:  Status(fmt.Errorf("load: %w", ErrNotFound))
Output: 404
```

**Example 3:**

```
Input:  Status(errors.New("boom"))
Output: 500
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Ordered classification** | The first matching sentinel wins. |
| 2 | **errors.Is over a chain** | Wrapped sentinels still classify. |
| 3 | **Default branch** | Unknown failures must map to something safe. |

## Hint

A `switch` with no expression reads better than four chained `if`s — but the nil case comes first.

## Validate

```bash
make verify
```
