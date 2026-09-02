# Type Switch

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A config loader accepts values of several dynamic types and renders each one for a text config file.

## Task

Implement the stub(s) in [typeswitch.go](typeswitch.go):

1. Implement `Render`, which formats an `any` value: `int` as digits, `bool` as `true`/`false`, `string` unchanged, `[]string` joined with `,`, and anything else as `"?"`.
2. Use a single type switch — no chain of separate assertions.
3. Keep it one pass: do not build intermediate slices for the `[]string` case beyond the output.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Render(42)
Output: "42"
```

**Example 2:**

```
Input:  Render([]string{"a", "b"})
Output: "a,b"
```

**Example 3:**

```
Input:  Render(3.5)
Output: "?"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type switch** | `switch v := x.(type)` binds the concrete value in each case. |
| 2 | **Interface dynamic type** | An `any` holds a (type, value) pair; the switch reads the type half. |
| 3 | **strconv + strings** | Reused from standard library: `Itoa` and `Join`. |

## Hint

`case []string:` binds `v` as a `[]string` — `strings.Join(v, ",")` works directly.

## Validate

```bash
make verify
```
