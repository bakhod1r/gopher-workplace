# A Byte View Of One Value

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A checksum routine takes `[]byte` and the caller has a counter. Encoding the counter into a temporary buffer allocates on every update.

## Task

Implement [valuebytes.go](valuebytes.go):

1. Return an 8-byte view of the value `p` points at, sharing its storage.
2. A nil pointer yields nil.
3. The view's length and capacity must both come from `unsafe.Sizeof` — zero allocations.

Replace the stub body in [valuebytes.go](valuebytes.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  v := uint64(0); b := Bytes(&v); v = ^uint64(0)
Output: b reads all 0xff
```

_Explanation:_ The view is live.

**Example 2:**

```
Input:  Bytes(nil)
Output: <nil>
```

**Example 3:**

```
Input:  cap of the result
Output: 8
```

_Explanation:_ So an append cannot write past the value.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **unsafe.Slice over a typed pointer** | Reinterprets one value as its bytes. |
| 2 | **Sizeof for the length** | Hard-coding 8 breaks the moment the type changes. |
| 3 | **Machine layout, not a wire format** | The byte order is the host's. |

## Hint

Convert the pointer to `*byte`, and take the length from the type.

## Validate

```bash
make verify
```
