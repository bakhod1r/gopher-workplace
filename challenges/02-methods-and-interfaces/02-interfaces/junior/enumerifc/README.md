# Enumerable Interface

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A report engine walks any collection that can hand out its items one index at a time.

## Task

Implement the stub(s) in [enumerifc.go](enumerifc.go):

1. Implement `Len` and `At` on `Words`.
2. Implement `Join`, which concatenates every item with the given separator.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Words{"a", "b"}.Len()
Output: 2
```

**Example 2:**

```
Input:  Words{"a", "b"}.At(1)
Output: "b"
```

**Example 3:**

```
Input:  Join(Words{"a", "b", "c"}, "-")
Output: "a-b-c"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Multi-method interface** | `Enumerable` needs both `Len` and `At` — a type must supply all of them. |
| 2 | **Defined slice type** | `type Words []string` carries methods. |
| 3 | **Index loop** | Reused: `for i := 0; i < n; i++` over positions. |

## Hint

Add the separator before every item except the first.

## Validate

```bash
make verify
```
