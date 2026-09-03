# Draining A Reader Into The Caller's Buffer

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

`io.ReadAll` allocates a fresh slice per call, which is fine once and expensive in a request loop. The reusable version takes the caller's buffer and grows it only when the payload is bigger than anything seen before — the same trick `bufio` uses, and the reason a server can serve millions of requests without touching the allocator.

## Task

Implement both functions in [chunkedreader.go](chunkedreader.go):

1. `ReadAll` fills `dst` (resetting it first) by reading in `chunk`-sized steps, reusing its capacity and growing only as needed.
2. A non-positive `chunk` uses 4096; `io.EOF` ends the read normally, any other error is returned *with* the data read so far.
3. `CountChunks` reports how many `Read` calls draining `n` bytes takes, including the final EOF call.

## Examples

**Example 1:**

```
Input:  ReadAll("hello", nil, 2)
Output: "hello", nil
```

**Example 2:**

```
Input:  ReadAll("new", []byte("stale"), 4)
Output: "new"
```

**Example 3:**

```
Input:  CountChunks(5, 2)
Output: 4
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`Read` may return less than asked** | A short read is normal, not an error; loop until EOF. |
| 2 | **EOF is not a failure** | It is the termination condition, and must never reach the caller as an error. |
| 3 | **Partial data on error** | Returning what was read lets the caller log or salvage it. |

## Topics used again

`io.Reader`, `errors.Is`, slice growth, `append`.

## Hint

Grow with `append(dst, make([]byte, chunk)...)`-style reslicing, or read into `dst[len:cap]` and extend the length by `n`.

## Validate

```bash
make verify
```
