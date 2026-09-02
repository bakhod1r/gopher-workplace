# Stringer and Formatting

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A CLI prints domain values, and the types control exactly how they appear in log lines.

## Task

Implement the stub(s) in [fmtstring.go](fmtstring.go):

1. Implement `String` on `Money` — `"<whole>.<cents padded to 2>"` from a value in cents.
2. Implement `String` on `Level` — `"DEBUG"`, `"INFO"`, `"ERROR"`, or `"LEVEL(n)"` for anything else.
3. Implement `Line`, which renders `"[<level>] <msg>: <money>"` using `fmt`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Money(1234).String()
Output: "12.34"
```

**Example 2:**

```
Input:  Money(5).String()
Output: "0.05"
```

**Example 3:**

```
Input:  Line(Info, "paid", Money(1234))
Output: "[INFO] paid: 12.34"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **fmt.Stringer** | `%v` and `%s` call `String()` automatically. |
| 2 | **Zero padding** | Reused from standard library: `%02d`. |
| 3 | **Negative handling** | Reused: sign is separate from the digits. |

## Hint

`fmt.Sprintf("%d.%02d", cents/100, cents%100)` for non-negative values.

## Validate

```bash
make verify
```
