# Bytes Are Not Characters

**Level:** junior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A field is limited to twenty characters. The check uses `len(s)`, and users with accented names are rejected for a name that fits.

## Task

Implement [bytesandrunes.go](bytesandrunes.go):

1. Return the byte length and the character count of `s`.
2. Zero allocations — do not convert to `[]rune`.

Replace the stub body in [bytesandrunes.go](bytesandrunes.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Counts("hello")
Output: 5, 5
```

_Explanation:_ ASCII: one byte per character.

**Example 2:**

```
Input:  Counts("héllo")
Output: 6, 5
```

_Explanation:_ é is two bytes.

**Example 3:**

```
Input:  Counts("日本語")
Output: 9, 3
```

_Explanation:_ Three bytes each.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **len is bytes** | A string's length is its storage size, not its character count. |
| 2 | **range decodes UTF-8** | Ranging a string yields one iteration per character. |
| 3 | **[]rune(s) allocates** | It builds a whole new slice of 32-bit values. |

## Hint

`range` over the string, and count the iterations.

## Validate

```bash
make verify
```
