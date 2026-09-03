# Split A Joined Error

**Level:** middle
**Topic:** 04-error-handling

## Context

A report lists every failed rule separately. The joined error it received must be taken apart again.

## Task

Implement `Split` in [joinsplit.go](joinsplit.go):

1. Return the individual errors inside an `errors.Join` result.
2. Return a single-element slice for an error that is not a join.
3. Return nil for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Split(errors.Join(ErrA, ErrB))
Output: [ErrA ErrB]
```

**Example 2:**

```
Input:  Split(ErrA)
Output: [ErrA]
```

**Example 3:**

```
Input:  Split(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Unwrap() []error** | Joined errors expose a multi-error Unwrap. |
| 2 | **Interface assertion on a method** | Assert to `interface{ Unwrap() []error }`. |
| 3 | **Two Unwrap shapes** | Single-error and multi-error unwrapping are different methods. |

## Hint

`errors.Unwrap` only knows the single-error form; a joined error implements the slice form instead.

## Validate

```bash
make verify
```
