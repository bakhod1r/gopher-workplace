# Strings and Bytes

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

A string and a `[]byte` are convertible both ways with a plain type conversion.
Each conversion copies the data.

## Task

Implement `FromBytes(b)` (bytes → string) and `ToBytes(s)` (string → bytes).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FromBytes([]byte{'G','o'})
Output: "Go"
```

_Explanation:_ string([]byte) copies the bytes into a string.

**Example 2:**

```
Input:  FromBytes(nil)
Output: ""
```

_Explanation:_ Converting a nil slice yields the empty string.

**Example 3:**

```
Input:  ToBytes("Go")
Output: [71 111]
```

_Explanation:_ []byte(string) copies the string's bytes.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **string([]byte)** | Converts bytes to a string (copy). |
| 2 | **[]byte(string)** | Converts a string to bytes (copy). |
| 3 | **Immutability** | Strings are read-only; the []byte is a mutable copy. |

## Hint

`string(b)` and `[]byte(s)`.

## Validate

```bash
make verify
```
