# The Interface That Is Not Nil

**Level:** senior
**Topic:** 10-advanced-topics / 03-reflection

## Context

A handler returns `err` from a helper that declares `*ValidationError` as its result type. Every call site sees a non-nil error, every request fails validation, and the error message is empty.

## Task

Implement [typednil.go](typednil.go):

1. Report whether `v` is nil or wraps a nil pointer, map, slice, channel, function or interface.
2. Values that cannot be nil — ints, strings, structs, arrays — report false.
3. Never panic, whatever the input.

Replace the stub body in [typednil.go](typednil.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  var p *int; IsNilValue(p)
Output: true
```

_Explanation:_ An interface holding a typed nil pointer.

**Example 2:**

```
Input:  IsNilValue([]int{})
Output: false
```

_Explanation:_ Empty is not nil.

**Example 3:**

```
Input:  IsNilValue(0)
Output: false
```

_Explanation:_ An int cannot be nil.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface representation** | An interface is a (type, value) pair; a typed nil has a type, so it is not == nil. |
| 2 | **Value.IsNil** | Valid only for the nilable kinds — it panics on the rest. |
| 3 | **Why the trap bites** | Assigning a typed nil pointer to an `error` variable makes it non-nil forever after. |

## Hint

`v == nil` is the first check, not the only one. And `IsNil` panics on an int.

## Validate

```bash
make verify
```
