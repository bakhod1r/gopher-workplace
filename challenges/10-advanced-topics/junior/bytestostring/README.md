# A String View Of Bytes You Own

**Level:** junior
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A parser converts every field to a string just to compare it. Each conversion copies bytes the parser already has and throws the copy away a line later.

## Task

Implement [bytestostring.go](bytestostring.go):

1. Return a string sharing `b`'s bytes, with no copy.
2. An empty or nil input returns the empty string.
3. Zero allocations, whatever the length.

Replace the stub body in [bytestostring.go](bytestostring.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Str([]byte("hello"))
Output: "hello"
```

**Example 2:**

```
Input:  Str(nil)
Output: ""
```

**Example 3:**

```
Input:  StringData(result) vs SliceData(input)
Output: the same pointer
```

_Explanation:_ No copy happened.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **unsafe.String** | Builds a string header over a pointer and a length. |
| 2 | **unsafe.SliceData** | The address of a slice's first element, defined even for empty slices. |
| 3 | **The obligation you take on** | The bytes must not change while the string is alive. |

## Hint

`unsafe.String` needs a `*byte` and a length. `unsafe.SliceData` provides the first.

## Validate

```bash
make verify
```
