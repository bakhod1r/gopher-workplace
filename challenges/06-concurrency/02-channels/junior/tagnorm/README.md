# Normalize Tags

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The metrics ingest path normalises label values before they reach storage:
a stage between the receiver and the writer upper-cases each tag and closes
its output when the receiver's stream ends.

## Task

Implement `NormalizeTags` in [tagnorm.go](tagnorm.go) so that:

1. It receives from `in` until `in` is closed.
2. For each tag it sends `strings.ToUpper(tag)` on `out`, in order.
3. It closes `out` once — it must never close `in`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NormalizeTags(in: "az","eu")
Output: out: "AZ", "EU" then closed
```

**Example 2:**

```
Input:  NormalizeTags(in: "Prod")
Output: out: "PROD" then closed
```

**Example 3:**

```
Input:  NormalizeTags(in: closed, empty)
Output: out: closed, no values
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Both directions** | `<-chan` in, `chan<- ` out — the shape of a pipeline stage. |
| 2 | **`range` + `close`** | Drain the receiver's stream, then close the writer's. |
| 3 | **`strings.ToUpper`** | Pure transformation applied per tag. |

## Hint

A pipeline stage is `for v := range in { out <- f(v) }` followed by
`close(out)`.

## Validate

```bash
make verify
```
