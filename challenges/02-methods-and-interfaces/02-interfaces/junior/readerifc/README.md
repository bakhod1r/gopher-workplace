# Reader Interface

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A parser pulls bytes from any source: a fixed string, or an empty source that is done immediately.

## Task

Implement the stub(s) in [readerifc.go](readerifc.go):

1. Implement `Read` on `*StringSource` — return the next chunk and whether anything was read.
2. Implement `ReadAll`, which drains the source and returns the concatenation.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  s := &StringSource{Data: "abc", Chunk: 2}; s.Read()
Output: "ab", true
```

**Example 2:**

```
Input:  second Read()
Output: "c", true; third Read() => "", false
```

**Example 3:**

```
Input:  ReadAll(&StringSource{Data: "hello", Chunk: 2})
Output: "hello"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Iterator-shaped interface** | `(value, ok)` is the Go idiom for a stream that can end. |
| 2 | **String slicing** | Reused from data types: `s[a:b]` and `len`. |
| 3 | **Pointer receiver state** | The read position must survive between calls. |

## Hint

Clamp the chunk end with `if end > len(s.Data) { end = len(s.Data) }`.

## Validate

```bash
make verify
```
