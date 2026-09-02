# Checksum Worker Pool

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The deduplication job checksums a directory through a worker pool. A bare
number coming back from a worker is useless — the caller cannot tell which
file it belongs to — so the results channel carries a small struct pairing the
file with its checksum.

## Task

Implement `ChecksumFiles` in [checksumpool.go](checksumpool.go) so that:

1. It creates a jobs channel and a buffered results channel of `result` structs, then starts `workers` goroutines.
2. Each worker ranges over jobs and sends a `result{file, sum(file)}`.
3. After `close(jobs)`, `wg.Wait()` and `close(results)`, the caller alone builds and returns the map.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ChecksumFiles([]string{"a"}, 2, sum)
Output: map[a:1]
```

**Example 2:**

```
Input:  ChecksumFiles([]string{"a", "bb"}, 2, sum)
Output: map[a:1 bb:2]
```

**Example 3:**

```
Input:  ChecksumFiles(nil, 3, sum)
Output: empty map
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Struct results** | Carry the key with the value so results stay attributable. |
| 2 | **Map built by one goroutine** | Only the caller writes the map, so no mutex is needed at all. |
| 3 | **Close ordering** | close(jobs) → wg.Wait() → close(results) → drain. |

## Hint

Send `result{file: file, sum: sum(file)}` from the workers, and populate the
map in the calling goroutine after the results channel is closed.

## Validate

```bash
make verify
```
