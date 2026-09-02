# cmp.Or

**Level:** junior  
**Topic:** 03-generics

## Context

A profile header shows the friendliest name available, falling back through a chain of options.

## Task

Implement the stub(s) in [cmporgen.go](cmporgen.go):

1. Implement `Display` using `cmp.Or`.
2. Return `"anonymous"` when both arguments are empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Display("nick", "user")
Output: "nick"
```

**Example 2:**

```
Input:  Display("", "user")
Output: "user"
```

**Example 3:**

```
Input:  Display("", "")
Output: "anonymous"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Or`** | Returns the first non-zero argument, or the zero value when all are zero. |
| 2 | **Zero means absent** | Reused from earlier: this is the same trade-off as your hand-written `Coalesce`. |
| 3 | **Variadic fallbacks** | Extra fallbacks are extra arguments — no nesting required. |

## Hint

One call: `cmp.Or(nickname, username, "anonymous")`.

## Validate

```bash
make verify
```
