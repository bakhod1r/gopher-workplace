# Compare Bytes To A String Without Converting

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A hot dispatcher compares an incoming frame's type field against a handful of known names. Every comparison copies the frame's bytes into a string first.

## Task

Implement [bytesequal.go](bytesequal.go):

1. Report whether `b`'s bytes equal `s`.
2. Convert nothing — zero allocations for any length.
3. Compare the lengths before the contents.

Replace the stub body in [bytesequal.go](bytesequal.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  EqualString([]byte("hi"), "hi")
Output: true
```

**Example 2:**

```
Input:  EqualString([]byte("his"), "hi")
Output: false
```

_Explanation:_ Different lengths cannot match.

**Example 3:**

```
Input:  EqualString(nil, "")
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **A borrowed view is safe here** | The string dies inside the call, so nothing can observe the aliasing. |
| 2 | **Length first** | It is one comparison and it settles most cases. |
| 3 | **bytes.Equal is the real answer** | This puzzle is about why it can be allocation-free. |

## Hint

Wrap the bytes in a string that lives only long enough to be compared.

## Validate

```bash
make verify
```
