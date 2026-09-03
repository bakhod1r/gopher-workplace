# Read A String's Bytes Without Copying

**Level:** junior
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A checksum routine takes `[]byte` and every caller has a string. `[]byte(s)` copies the whole payload just so it can be read once.

## Task

Implement [stringtobytes.go](stringtobytes.go):

1. Return a byte view of `s` without copying.
2. The empty string yields an empty result.
3. Zero allocations; the result's capacity must equal its length.

Replace the stub body in [stringtobytes.go](stringtobytes.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Bytes("hello")
Output: []byte("hello")
```

_Explanation:_ Sharing the string's bytes.

**Example 2:**

```
Input:  Bytes("")
Output: []
```

**Example 3:**

```
Input:  cap(Bytes("abcd"))
Output: 4
```

_Explanation:_ So an append cannot write past the string.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **unsafe.Slice** | Builds a slice header over a pointer and a length. |
| 2 | **unsafe.StringData** | The address of a string's first byte. |
| 3 | **Read-only in practice** | String literals may live in a read-only mapping; writing there faults. |

## Hint

The mirror image of the byte-to-string direction.

## Validate

```bash
make verify
```
