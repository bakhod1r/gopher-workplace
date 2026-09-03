# One Allocation For A Stream Of Unknown Length

**Level:** senior
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

An RPC layer knows the payload size from the frame header but still calls `io.ReadAll`, which discovers the size by doubling — five allocations and four copies for a payload it could have sized exactly.

## Task

Implement [growonce.go](growonce.go):

1. Read `r` to EOF and return its bytes.
2. With an accurate `hint`, use at most two allocations.
3. An under- or over-estimated hint must still return the correct bytes; read errors propagate.

Replace the stub body in [growonce.go](growonce.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Collect(strings.NewReader("abc"), 3)
Output: "abc", nil
```

**Example 2:**

```
Input:  Collect(strings.NewReader(strings.Repeat("x",5000)), 4)
Output: all 5000 bytes
```

_Explanation:_ A bad hint costs speed, not correctness.

**Example 3:**

```
Input:  a failing reader
Output: the error
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Read into spare capacity** | `r.Read(buf[len(buf):cap(buf)])` fills the room you already own. |
| 2 | **append as the growth policy** | Appending one byte and reslicing back is the idiomatic "grow if full". |
| 3 | **Partial reads with EOF** | `Read` may return bytes and `io.EOF` together. |

## Hint

Read into `buf[len(buf):cap(buf)]` and only grow when that window is empty.

## Validate

```bash
make verify
```
