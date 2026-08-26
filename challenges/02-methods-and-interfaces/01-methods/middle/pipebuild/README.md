# Pipe Builder

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Method chaining is often used for string manipulation pipelines, applying
transformations step by step.

## Task

Implement `Upper` and `Replace` on `*Pipe` in [pipebuild.go](pipebuild.go):

1. `Upper`: mutate `text` using `strings.ToUpper`, return `p`.
2. `Replace`: mutate `text` using `strings.ReplaceAll`, return `p`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  NewPipe("hello").Upper().Result()
Output: "HELLO"
```

**Example 2:**

```
Input:  NewPipe("go lang").Upper().Replace(" ", "-").Result()
Output: "GO-LANG"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Chaining mutations** | Each method updates state and returns the pointer. |
| 2 | **strings package** | Common string ops. |

## Hint

`p.text = strings.ToUpper(p.text); return p`.

## Validate

```bash
make verify
```
