# The Pooled Reader Still Reading The Last Request

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A request handler pools its `bufio.Reader` to cut allocations. Under load, responses start containing fragments of other requests' bodies.

## Task

Fix the single planted bug in [pooledreader.go](pooledreader.go):

1. Return the first line of `r`, without the trailing newline.
2. A final line with no newline still counts; an empty reader returns the empty string.
3. Fix the single bug so a pooled reader reads from the right source.
4. The reader must still go back to the pool.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  FirstLine(strings.NewReader("alpha\nbeta\n"))
Output: "alpha", nil
```

**Example 2:**

```
Input:  100 sequential requests
Output: each gets its own first line
```

**Example 3:**

```
Input:  a failing reader
Output: the error
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pooled values carry state** | A `bufio.Reader` holds a source and a buffer of unread bytes. |
| 2 | **Reader.Reset** | Rebinds the reader to a new source and discards the buffer. |
| 3 | **Cross-request contamination** | Reuse bugs leak one caller's data into another's output. |
| 4 | **io.EOF is not a failure** | A final unterminated line arrives with EOF. |

## Hint

The pool hands you a reader. What is it currently reading from?

## Validate

```bash
make verify
```
