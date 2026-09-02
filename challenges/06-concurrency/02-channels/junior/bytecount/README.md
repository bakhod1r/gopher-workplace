# Total Bytes

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

A CDN edge node reports bandwidth per point of presence. Each served
response pushes its body size onto a channel; when the reporting window
ends the producer closes the channel and the accountant emits the total.

## Task

Implement `TotalBytes` in [bytecount.go](bytecount.go) so that:

1. It receives sizes until `sizes` is closed.
2. It returns the sum of everything received.
3. A window with no responses (an already-closed channel) returns `0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TotalBytes(1200, 800)
Output: 2000
```

**Example 2:**

```
Input:  TotalBytes(512)
Output: 512
```

**Example 3:**

```
Input:  TotalBytes() // closed, empty
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`range` over a channel** | Loops until the channel is closed and drained. |
| 2 | **Receive-only parameter** | `<-chan int` documents that the accountant never sends or closes. |
| 3 | **Accumulator** | A plain local starting at the identity element `0`. |

## Hint

`for n := range sizes` ends by itself when the edge closes the channel —
no counter or comma-ok needed.

## Validate

```bash
make verify
```
