# Let The Caller Own The Buffer

**Level:** junior
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A codec allocates and returns a fresh buffer on every call. The caller has a perfectly good buffer it would rather reuse, but the API gives it no way to say so.

## Task

Implement [outparam.go](outparam.go):

1. Write `v` into every byte of `dst`.
2. Return the number of bytes written.
3. Allocate nothing — the storage is the caller's.

Replace the stub body in [outparam.go](outparam.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Fill(make([]byte,3), 'x')
Output: 3, buffer "xxx"
```

**Example 2:**

```
Input:  buf := []byte("abcd"); Fill(buf[1:3], 'z')
Output: "azzd"
```

_Explanation:_ Only the view is written.

**Example 3:**

```
Input:  Fill(nil, 'x')
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Caller-owned buffers** | A `dst` parameter moves the allocation decision to the caller. |
| 2 | **Nothing escapes** | The function keeps no reference after it returns. |
| 3 | **Views** | A sub-slice limits exactly which bytes may be touched. |

## Hint

You are given the memory. Do not ask for more.

## Validate

```bash
make verify
```
