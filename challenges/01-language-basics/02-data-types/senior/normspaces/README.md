# Normalize Whitespace

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A search indexer normalizes whitespace so `"  hello   world  "` becomes
`"hello world"`. The current code prepends a stray space, so keys never match.

## Task

Fix the return between the markers in [normspaces.go](normspaces.go). Note
`strings.Fields` already dropped the empties — no manual space needed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "  hello   world  "
Output: "hello world"
```

**Example 2:**

```
Input:  "a\tb\nc"
Output: "a b c"
```

**Example 3:**

```
Input:  ""
Output: ""
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Fields + Join** | `Fields` splits on whitespace runs; `Join` re-glues. |
| 2 | **No extra space** | The join already omits leading/trailing. |
| 3 | **All whitespace kinds** | Tabs/newlines collapse too. |

## Hint

`return collapsed`.

## Validate

```bash
make verify
```
