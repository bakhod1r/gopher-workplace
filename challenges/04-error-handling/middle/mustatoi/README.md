# Wrap A Library Error

**Level:** middle
**Topic:** 04-error-handling

## Context

A CSV column is parsed as a number. When it fails, the report must show the offending text alongside the library's own explanation.

## Task

Implement `ParseInt` in [mustatoi.go](mustatoi.go):

1. Return the parsed number and nil on success.
2. Return `0` and an error reading `"parse \"<s>\": <cause>"` on failure.
3. Keep `strconv.ErrSyntax` matchable through the wrapper.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ParseInt("42")
Output: 42, nil
```

**Example 2:**

```
Input:  ParseInt("x")
Output: 0, "parse \"x\": …"
```

**Example 3:**

```
Input:  errors.Is(err, strconv.ErrSyntax)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Wrapping third-party errors** | Add context without discarding the cause. |
| 2 | **Quoting input with %q** | The offending text is shown unambiguously. |
| 3 | **Standard library sentinels** | `strconv.ErrSyntax` is matchable. |

## Hint

`strconv.Atoi` already returns a rich `*strconv.NumError` — wrap it, do not replace it.

## Validate

```bash
make verify
```
