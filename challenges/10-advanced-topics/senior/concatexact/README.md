# Join With Exactly One Allocation

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A hot path joins five short strings per request. `strings.Join` copies once into its buffer and `string(buf)` copies again — the second copy is pure waste on a buffer nothing else can see.

## Task

Implement [concatexact.go](concatexact.go):

1. Return the parts joined end to end.
2. Exactly one allocation: size the buffer from the total length, then hand it out without copying again.
3. An empty input, or all-empty parts, returns the empty string.

Replace the stub body in [concatexact.go](concatexact.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Concat([]string{"a","bc","d"})
Output: "abcd"
```

**Example 2:**

```
Input:  Concat(nil)
Output: ""
```

**Example 3:**

```
Input:  allocations per call
Output: 1
```

_Explanation:_ The buffer, and nothing else.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Exact sizing** | Summing the lengths removes every growth step. |
| 2 | **unsafe.String over your own buffer** | Legal precisely because nothing else can write to it. |
| 3 | **string(buf) is the second copy** | It is required when the buffer is shared, and wasted when it is not. |
| 4 | **Local ownership** | The buffer never leaves the function except as the string. |

## Hint

Build it in a `[]byte` you allocated, then stop copying.

## Validate

```bash
make verify
```
