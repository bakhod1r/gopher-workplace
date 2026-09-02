# The Default Branch That Eats The Value

**Level:** staff  
**Topic:** 03-generics

## Context

An ingest pipeline normalises every field before writing it. String fields survive; every numeric and boolean field arrives at the database as 0 or false, and nobody noticed because the writes succeed.

## Task

Fix the single planted bug in [typeswitchdefaultbug.go](typeswitchdefaultbug.go):

1. Find and fix the single bug so a type the switch does not special-case is returned unchanged.
2. String normalisation must keep working.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Normalize(" Hi ")
Output: "hi"
```

**Example 2:**

```
Input:  Normalize(42)
Output: 42
```

**Example 3:**

```
Input:  Normalize(true)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Generics are not `any`** | A type parameter is resolved at compile time — routing values through `interface{}` throws that away. |
| 2 | **The default branch is the common case** | Most values fall through it, so a wrong default is a wrong function. |
| 3 | **Zero values are not "no answer"** | `var zero T` is a valid-looking result that silently replaces real data. |

## Hint

What does the switch return for an `int`?

## Validate

```bash
make verify
```
