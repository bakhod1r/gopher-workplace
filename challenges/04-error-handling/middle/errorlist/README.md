# Aggregate Error Type

**Level:** middle
**Topic:** 04-error-handling

## Context

A schema check collects every violation and returns them as one error whose message lists them all.

## Task

Implement `Errors` in [errorlist.go](errorlist.go):

1. Give `Errors` an `Error() string` joining the messages with `"; "`.
2. Give it an `Unwrap() []error` so `errors.Is` finds each member.
3. An empty `Errors` renders as the empty string.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Errors{ErrA, ErrB}.Error()
Output: "rule a; rule b"
```

**Example 2:**

```
Input:  errors.Is(Errors{ErrA}, ErrA)
Output: true
```

**Example 3:**

```
Input:  Errors{}.Error()
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Slice types with methods** | `type Errors []error` can implement `error`. |
| 2 | **Multi-error Unwrap** | `Unwrap() []error` plugs into `errors.Is`. |
| 3 | **strings.Join** | Building the combined message. |

## Hint

Two methods are missing. `errors.Is` never looks at your message — it uses the other one.

## Validate

```bash
make verify
```
