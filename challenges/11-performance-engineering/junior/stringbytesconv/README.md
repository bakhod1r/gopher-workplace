# The Conversion You Do Not Need

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

`[]byte(s)` copies the entire string, because strings are immutable and byte slices are not. In a hot loop that copy is often the only allocation in the function — and it is usually there for no reason, since a string can be indexed and ranged directly.

## Task

Implement both functions in [stringbytesconv.go](stringbytesconv.go):

1. `CountByte` returns how many times `b` occurs in `s`, counting bytes, not runes.
2. `HasPrefixByte` reports whether `s` starts with `b`; the empty string starts with nothing.
3. Neither function may allocate.

## Examples

**Example 1:**

```
Input:  CountByte("hello", 'l')
Output: 2
```

**Example 2:**

```
Input:  CountByte("é", 0xC3)
Output: 1
```

**Example 3:**

```
Input:  HasPrefixByte("", 'a')
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`[]byte(s)` copies** | The cost is proportional to the string, and it shows up as one allocation per call. |
| 2 | **`s[i]` is a byte** | Indexing a string is direct memory access with no conversion. |
| 3 | **`for i := range` vs `for _, r := range`** | The two-variable form decodes UTF-8 runes; the index form walks bytes. |

## Topics used again

Strings, bytes, UTF-8 encoding, loops.

## Hint

Index with `s[i]`; never write `[]byte(s)`.

## Validate

```bash
make verify
```
