# Count The Lines Without Holding The File

**Level:** senior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A log tool calls `io.ReadAll` and then counts newlines. It works on the developer's sample and is killed by the OOM reaper on the first real file.

## Task

Implement [streamcount.go](streamcount.go):

1. Return the number of `\n` bytes in the stream.
2. Read through a fixed-size buffer — total allocation must stay under 1 MiB regardless of the stream size.
3. Return any read error other than `io.EOF` along with the count so far.

Replace the stub body in [streamcount.go](streamcount.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  CountLines(strings.NewReader("a\nb\nc"))
Output: 2, nil
```

_Explanation:_ The final unterminated line is not counted.

**Example 2:**

```
Input:  CountLines(strings.NewReader(""))
Output: 0, nil
```

**Example 3:**

```
Input:  a reader that fails
Output: the error
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Streaming vs buffering** | Memory should track the buffer size, not the input size. |
| 2 | **io.Reader contract** | `Read` may return `n > 0` together with an error; count those bytes first. |
| 3 | **io.EOF is not a failure** | It ends the loop and returns a nil error. |

## Hint

`io.ReadAll` is the bug, not the tool. What size is your working set allowed to be?

## Validate

```bash
make verify
```
