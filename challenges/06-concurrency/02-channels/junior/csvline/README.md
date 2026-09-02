# Join Fields

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The report exporter receives one record's field values on a channel and
assembles a single CSV line from them, separated by the configured
delimiter.

## Task

Implement `JoinFields` in [csvline.go](csvline.go) so that:

1. It drains `fields` until the record ends.
2. It joins the received fields with `sep` **between** them — no trailing delimiter.
3. An empty record returns `""`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  JoinFields("a","b" | sep ",")
Output: "a,b"
```

**Example 2:**

```
Input:  JoinFields("x" | sep ",")
Output: "x"
```

**Example 3:**

```
Input:  JoinFields(empty | sep ",")
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Collect then join** | `strings.Join` needs a slice, so drain first. |
| 2 | **Separator placement** | `strings.Join` puts `sep` between, never at the ends. |
| 3 | **`range` over a channel** | Ends when the record's fields are exhausted. |

## Hint

Collect into a `[]string` first, then hand it to `strings.Join` — it gets
the delimiter placement right for free.

## Validate

```bash
make verify
```
