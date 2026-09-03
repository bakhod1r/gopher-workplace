# Annotate With Position

**Level:** middle
**Topic:** 04-error-handling

## Context

A batch parser reports which record failed. The record number belongs in the message, but the underlying cause must stay matchable.

## Task

Implement `AtIndex` in [indexwrap.go](indexwrap.go):

1. Return nil when `err` is nil.
2. Return an error whose message is `"record <i>: <err>"`.
3. Keep the wrapped error reachable by `errors.Is`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  AtIndex(3, ErrParse)
Output: "record 3: parse failed"
```

**Example 2:**

```
Input:  AtIndex(0, nil)
Output: nil
```

**Example 3:**

```
Input:  errors.Is(AtIndex(3, ErrParse), ErrParse)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Formatting with %w** | `%d` and `%w` combine in one call. |
| 2 | **Verb order** | The wrapped error can appear anywhere in the format. |
| 3 | **Nil guard** | Annotation of nothing is still nothing. |

## Hint

One `fmt.Errorf` call carries both the index and the wrapped error.

## Validate

```bash
make verify
```
