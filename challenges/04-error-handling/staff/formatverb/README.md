# Verbose Formatting

**Level:** staff
**Topic:** 04-error-handling

## Context

Log lines print the short message, but a debug flag prints the full detail. One error type serves both through the formatting verb.

## Task

Implement `DetailError` in [formatverb.go](formatverb.go):

1. Give `*DetailError` an `Error() string` returning `Msg`.
2. Implement `Format` so `%v` and `%s` print `Msg`.
3. Make `%+v` print `"<Msg>\n\t<Detail>"`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  fmt.Sprintf("%v", e)
Output: "failed"
```

**Example 2:**

```
Input:  fmt.Sprintf("%+v", e)
Output: "failed\n\tat line 4"
```

**Example 3:**

```
Input:  fmt.Sprintf("%s", e)
Output: "failed"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **fmt.Formatter** | One method controls every verb. |
| 2 | **fmt.State** | `s.Flag('+')` reports the plus flag. |
| 3 | **Two audiences** | Terse by default, detailed on request. |

## Hint

`Format(s fmt.State, verb rune)` receives the verb and the flags — write to `s` with `io.WriteString`.

## Validate

```bash
make verify
```
