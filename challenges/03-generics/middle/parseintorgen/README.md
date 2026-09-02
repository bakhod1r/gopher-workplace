# Parse Or Default

**Level:** middle  
**Topic:** 03-generics

## Context

Configuration values arrive as strings. A malformed value should fall back to a default rather than stopping startup.

## Task

Implement the stub(s) in [parseintorgen.go](parseintorgen.go):

1. Implement `ParseOr`, returning the parsed value or `def`.
2. The result keeps the caller's integer type, including named types.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ParseOr[int]("42", 0)
Output: 42
```

**Example 2:**

```
Input:  ParseOr[int]("abc", 7)
Output: 7
```

**Example 3:**

```
Input:  ParseOr("5", Retries(1))
Output: Retries(5)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bridging stdlib to a type parameter** | `strconv` returns `int64`; `T(n)` narrows it to the caller's type. |
| 2 | **Inference from the default** | `ParseOr("5", Retries(1))` infers `T` from the second argument. |
| 3 | **Narrowing is unchecked** | Converting a large `int64` into a narrower `T` silently truncates. |

## Hint

`T(n)` converts the parsed `int64` into whatever the caller asked for.

## Validate

```bash
make verify
```
