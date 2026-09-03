# Classify By Table

**Level:** staff
**Topic:** 04-error-handling

## Context

A gateway maps dozens of sentinels onto status codes. A long `switch` grows unreadable, so the mapping lives in data.

## Task

Implement `Status` in [sentineltable.go](sentineltable.go):

1. Return the code registered for the first sentinel matching `err`.
2. Check sentinels in the order they appear in the table.
3. Return `500` for an unmatched error and `200` for nil.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Status(ErrNotFound, table)
Output: 404
```

**Example 2:**

```
Input:  Status(fmt.Errorf("x: %w", ErrDenied), table)
Output: 403
```

**Example 3:**

```
Input:  Status(errors.New("x"), table)
Output: 500
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Data-driven mapping** | The table is the policy. |
| 2 | **Ordered rules** | The first match wins, so order is meaningful. |
| 3 | **errors.Is per entry** | Wrapped sentinels still classify. |

## Hint

A slice preserves order in a way a map cannot — that is why the table is not a `map[error]int`.

## Validate

```bash
make verify
```
