# The Method That Copies Its Receiver

**Level:** middle
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A config struct grows a 512-byte field. Every accessor on it still takes a value receiver, and a benchmark that never touched the allocator starts copying half a kilobyte per call.

## Task

Implement [valuereceiver.go](valuereceiver.go):

1. Return the read and write timeouts from the receiver.
2. The receiver must be the caller's `Config`, so later writes are visible.
3. Zero allocations, zero large copies.

Replace the stub body in [valuereceiver.go](valuereceiver.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  (&Config{Read:5, Write:9}).Timeouts()
Output: 5, 9
```

**Example 2:**

```
Input:  c.Read = 42; c.Timeouts()
Output: 42, ...
```

_Explanation:_ The method sees the caller's struct.

**Example 3:**

```
Input:  (&Config{}).Timeouts()
Output: 0, 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer vs value receivers** | A value receiver copies the struct on every call. |
| 2 | **Copy cost scales with the struct** | Two words is free; 528 bytes is not. |
| 3 | **Receiver consistency** | Mixing receiver kinds on one type is a readability trap. |

## Hint

The signature is already given. The body is two field reads.

## Validate

```bash
make verify
```
