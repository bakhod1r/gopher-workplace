# io.Writer

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A report writer targets any `io.Writer`: a buffer in tests, a file in production.

## Task

Implement the stub(s) in [iowriter.go](iowriter.go):

1. Implement `WriteReport`, which writes `"title\n"` followed by one `"- <item>\n"` per item, and returns the total bytes written.
2. Stop and return the underlying error as soon as a write fails.
3. Do not build the whole report in one big string first — write incrementally.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  WriteReport(&bytes.Buffer{}, "T", []string{"a"})
Output: 6, nil (buffer holds "T\n- a\n")
```

**Example 2:**

```
Input:  WriteReport(&bytes.Buffer{}, "T", nil)
Output: 2, nil
```

**Example 3:**

```
Input:  a failing writer
Output: bytes written so far, the error
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **io.Writer** | The standard sink interface: `Write([]byte) (int, error)`. |
| 2 | **Byte counting and errors** | Every write returns both; both must be handled. |
| 3 | **Incremental output** | Efficiency: avoid materialising the whole report in memory. |

## Hint

`io.WriteString(w, s)` returns `(n, err)` — accumulate `n`, return on the first `err`.

## Validate

```bash
make verify
```
