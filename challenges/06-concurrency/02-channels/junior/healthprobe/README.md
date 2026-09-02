# Health Probe

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

A service-mesh sidecar validates a new upstream connection by exchanging
probe/ack round trips with it before marking the endpoint healthy. Each
round trip must complete before the next one starts.

## Task

Implement `Probe` in [healthprobe.go](healthprobe.go) so that:

1. It runs `rounds` round trips, tracing `"probe"` then `"ack"` for each.
2. The trace has `2 * rounds` entries in strict alternation.
3. `rounds <= 0` returns an empty, non-nil trace.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Probe(1)
Output: ["probe" "ack"]
```

**Example 2:**

```
Input:  Probe(2)
Output: ["probe" "ack" "probe" "ack"]
```

**Example 3:**

```
Input:  Probe(0)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Unbuffered channel** | A send blocks until a receive happens — a rendezvous. |
| 2 | **Goroutine** | `go func(){...}()` runs the upstream responder concurrently. |
| 3 | **Deterministic handoff** | Strict alternation removes all nondeterminism from the trace. |

## Hint

Two unbuffered channels: the sidecar sends on `probes` and receives on
`acks`; the upstream does the mirror image.

## Validate

```bash
make verify
```
