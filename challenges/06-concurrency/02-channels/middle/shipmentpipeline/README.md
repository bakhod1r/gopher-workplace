# Shipment Label Pipeline

**Level:** middle
**Topic:** 06-concurrency → 02-channels

## Context

A fulfilment service turns raw order IDs into printed shipping labels in three stages: validate, price, and render. Each stage is a goroutine joined by channels, and a stage that rejects an order must drop it so the later stages never see it.

## Task

Implement the stubbed function in [shipmentpipeline.go](shipmentpipeline.go) so that:

1. Feed the orders into the pipeline from their own goroutine.
2. Run the `keep` filter as one stage and `render` as the next, each in its own goroutine joined by channels.
3. Drop rejected orders before the render stage sees them.
4. Return the labels in input order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Labels([]string{"o1","o2"}, all, label)
Output: ["LABEL-o1" "LABEL-o2"]
```

**Example 2:**

```
Input:  Labels([]string{"o1","bad"}, notBad, label)
Output: ["LABEL-o1"]
```

**Example 3:**

```
Input:  Labels(nil, all, label)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | Pipeline stage | Each stage: range the input channel, write to an output channel, close the output when done. |
| 2 | Close propagation | Closing the source ends stage one's `range`, whose `defer close` ends stage two — shutdown flows downstream on its own. |
| 3 | Order preservation | A single goroutine per stage keeps the sequence; parallelism inside a stage would not. |
| 4 | Filtering stage | A stage may emit fewer items than it reads — that is how work is dropped. |

## Hint

Three goroutines, three channels, and `defer close(out)` in every one of them. The caller drains the last channel.

## Validate

```bash
make verify
```
