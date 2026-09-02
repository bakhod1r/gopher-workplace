# First Replica Ack

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

A quorum write fans out to every replica and only needs the first success to
return to the client. The subtlety is the exit: the goroutines for the slower
replicas are still running, and if the channel they report on has no room for
their answer they block on that send forever.

## Task

Implement `FirstReplicaAck` in [replicawrite.go](replicawrite.go) so that:

1. With no replicas it reports false immediately.
2. It starts one goroutine per replica, each sending its `write` result on a channel buffered to `len(replicas)`.
3. It reads results and returns true on the first ack, or false after every replica has answered.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FirstReplicaAck([]string{"ok-1", "bad-2"}, write)
Output: true
```

**Example 2:**

```
Input:  FirstReplicaAck([]string{"bad-1", "bad-2"}, write)
Output: false
```

**Example 3:**

```
Input:  FirstReplicaAck(nil, write)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fan-out with early return** | The caller stops reading before the workers stop working. |
| 2 | **Buffer sized to the workers** | Enough room for every result means no goroutine can block on send. |
| 3 | **Goroutine leak** | An unbuffered channel plus an early return strands every remaining writer. |

## Hint

Size the buffer to `len(replicas)`. Then every goroutine can deposit its
answer and exit even though the caller has already returned.

## Validate

```bash
make verify
```
