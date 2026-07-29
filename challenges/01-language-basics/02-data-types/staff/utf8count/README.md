# Rune Count by Lead Bytes

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A metrics job counts characters without allocating a `[]rune`. The rule is "count
bytes that are not continuation bytes", but the condition is inverted — it counts
the continuation bytes, so ASCII counts as 0 and multi-byte text is undercounted.

## Task

Fix the condition between the markers in [utf8count.go](utf8count.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "hello"
Output: 5
```

**Example 2:**

```
Input:  "héllo"
Output: 5
```

**Example 3:**

```
Input:  "日本語"
Output: 3
```

**Example 4:**

```
Input:  "a€b"
Output: 3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Continuation byte** | `c&0xC0 == 0x80` is `10xxxxxx`. |
| 2 | **One lead per rune** | Count the non-continuation bytes. |
| 3 | **No allocation** | Counts in place, unlike `[]rune`. |

## Hint

`if c&0xC0 != 0x80`.

## Validate

```bash
make verify
```
