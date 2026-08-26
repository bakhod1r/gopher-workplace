# Full Name

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A user profile page displays a person's full name by joining first and last
names with a space.

## Task

Implement `FullName` on `Person` in [fullname.go](fullname.go):

1. Return `First + " " + Last`.
2. Handle empty names — always include the space separator.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Person{"Go", "Gopher"}.FullName()
Output: "Go Gopher"
```

**Example 2:**

```
Input:  Person{"", "Doe"}.FullName()
Output: " Doe"
```

**Example 3:**

```
Input:  Person{"Jane", ""}.FullName()
Output: "Jane "
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Methods vs functions** | `p.FullName()` is a method tied to `Person`. |
| 2 | **Value receiver** | Read-only — no mutation needed. |
| 3 | **String concatenation** | Join with `+` or `fmt.Sprintf`. |

## Hint

`p.First + " " + p.Last` — keep the space unconditional.

## Validate

```bash
make verify
```
