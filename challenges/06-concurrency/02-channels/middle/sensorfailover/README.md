# Fail Over Between Sensor Gateways

**Level:** middle
**Topic:** 06-concurrency → 02-channels

## Context

The line controller reads temperature readings from two gateways at once. Each
gateway closes its stream when it finishes its shift, and one usually finishes
long before the other. A closed channel is *permanently ready to receive*, so a
naive two-arm `select` will spin on the finished gateway and never let the live
one through.

## Task

Implement `MergeSensorStreams` in [sensorfailover.go](sensorfailover.go) so that:

1. It loops while either stream is still live, selecting on both.
2. It uses comma-ok on each receive; on a closed stream it sets that local variable to `nil` and continues.
3. A `nil` channel arm blocks forever, which removes the finished gateway from the `select`.
4. It returns every reading collected, as a non-nil slice; a `nil` gateway argument counts as already finished.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MergeSensorStreams(primary[t1 t2], backup[t3])
Output: 3 readings (t1, t2, t3 in some order)
```

**Example 2:**

```
Input:  MergeSensorStreams(primary[t1 t2], backup=nil)
Output: 2 readings
```

**Example 3:**

```
Input:  MergeSensorStreams(closed empty, closed empty)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nil channel in select** | A nil arm never becomes ready — the idiom for disabling a case. |
| 2 | **Closed channel is always ready** | It yields the zero value instantly, forever. |
| 3 | **Comma-ok** | The only way a `select` receive can tell "value" from "closed". |
| 4 | **Loop condition as liveness** | `for primary != nil \|\| backup != nil` says "while anyone can still speak". |

## Hint

The parameters are just local variables. Assigning `primary = nil` inside the
function disables that arm without touching the caller's channel.

## Validate

```bash
make verify
```
