# Hand Out Bytes That Cannot Be Written

**Level:** senior
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A helper is described as "the fast way to get bytes from a string". A caller sorts the result in place, and a string constant elsewhere in the binary comes back reordered — when the process does not simply fault.

## Task

Fix the single planted bug in [freezebytes.go](freezebytes.go):

1. Return a byte slice holding `s`'s bytes that the caller may modify.
2. The empty string yields an empty result.
3. Fix the single bug: the result must not alias the string.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  b := Snapshot("abcd"); b[0] = 'X'
Output: s is unchanged
```

**Example 2:**

```
Input:  two snapshots of one string
Output: independent
```

**Example 3:**

```
Input:  Snapshot("")
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **String memory may be read-only** | Literals are placed in a read-only section; writing there faults. |
| 2 | **Immutability is program-wide** | Every holder of the string would observe a write through an aliased slice. |
| 3 | **The safe direction is asymmetric** | Bytes-to-string can be zero-copy under a no-write promise; string-to-writable-bytes cannot. |

## Hint

The conversion is legal. The promise the function makes is not.

## Validate

```bash
make verify
```
