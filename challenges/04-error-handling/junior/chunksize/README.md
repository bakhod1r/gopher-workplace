# Even Chunks

**Level:** junior
**Topic:** 04-error-handling

## Context

A batch uploader splits a total number of records into equally sized chunks. A zero or negative chunk count is a caller mistake.

## Task

Implement `ChunkSize` in [chunksize.go](chunksize.go):

1. Return `total / parts` rounded up, so no records are left over.
2. Return `0` and `ErrBadParts` when `parts <= 0`.
3. Return `0` and `ErrNegativeTotal` when `total < 0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ChunkSize(10, 3)
Output: 4, nil
```

**Example 2:**

```
Input:  ChunkSize(9, 3)
Output: 3, nil
```

**Example 3:**

```
Input:  ChunkSize(10, 0)
Output: 0, ErrBadParts
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Ceiling division** | `(a + b - 1) / b` rounds up with integers. |
| 2 | **Multiple guards** | Two independent invalid inputs. |
| 3 | **Integer truncation** | Plain `/` rounds down and loses records. |

## Hint

Plain integer division leaves a remainder unhandled — `10/3` is 3, and three chunks of 3 only cover 9 records.

## Validate

```bash
make verify
```
