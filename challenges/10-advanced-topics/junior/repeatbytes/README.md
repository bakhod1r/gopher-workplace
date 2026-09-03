# Repeat With The Length You Already Know

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A padding helper builds its filler by appending the pattern in a loop. For a 64 KB pad it reallocates seventeen times.

## Task

Implement [repeatbytes.go](repeatbytes.go):

1. Return `n` copies of `b` concatenated.
2. The result must share no storage with `b`.
3. One allocation only; `n <= 0` or an empty `b` gives an empty result.

Replace the stub body in [repeatbytes.go](repeatbytes.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Repeat([]byte("ab"), 3)
Output: "ababab"
```

**Example 2:**

```
Input:  Repeat([]byte("x"), 0)
Output: ""
```

**Example 3:**

```
Input:  Repeat(nil, 4)
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Known output size** | `len(b)*n` is the exact final length. |
| 2 | **copy into a window** | `copy(out[i*len(b):], b)` writes each repetition in place. |
| 3 | **Independence** | A fresh array means later writes to `b` are invisible. |

## Hint

Allocate `len(b)*n` bytes first; then it is `n` copies into windows.

## Validate

```bash
make verify
```
