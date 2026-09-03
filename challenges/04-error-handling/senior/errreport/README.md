# Render A Report

**Level:** senior
**Topic:** 04-error-handling

## Context

A CLI prints a failure summary: one numbered line per problem, in the order they were collected.

## Task

Implement `Report` in [errreport.go](errreport.go):

1. Return `""` for a nil error.
2. Return one line per joined member, formatted `"<n>. <message>"` starting at 1.
3. Treat a non-joined error as a single-line report.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Report(nil)
Output: ""
```

**Example 2:**

```
Input:  Report(ErrA)
Output: "1. a"
```

**Example 3:**

```
Input:  Report(errors.Join(ErrA, ErrB))
Output: "1. a\n2. b"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Presentation at the edge** | Rendering is separate from matching. |
| 2 | **Unwrap() []error** | The joined members drive the list. |
| 3 | **strings.Builder** | Assembling multi-line output. |

## Hint

The numbering starts at 1 while the slice index starts at 0.

## Validate

```bash
make verify
```
