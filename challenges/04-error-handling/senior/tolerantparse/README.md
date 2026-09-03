# Parse What You Can

**Level:** senior
**Topic:** 04-error-handling

## Context

An import accepts a partially broken file: valid rows are kept and every bad row is reported with its position.

## Task

Implement `Parse` in [tolerantparse.go](tolerantparse.go):

1. Return the successfully parsed values in input order.
2. Report each failure as `"line <i>: <cause>"`, joined into one error.
3. Return the parsed values even when some lines failed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Parse([]string{"1", "x", "3"})
Output: [1 3], one failure
```

**Example 2:**

```
Input:  Parse(nil)
Output: nil, nil
```

**Example 3:**

```
Input:  Parse([]string{"1"})
Output: [1], nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Partial success** | Results and failures are returned together. |
| 2 | **Index in the message** | Position is what makes a row fixable. |
| 3 | **Wrapping the library error** | `strconv.ErrSyntax` stays matchable. |

## Hint

Both return values matter on the failure path — the caller keeps the good rows.

## Validate

```bash
make verify
```
