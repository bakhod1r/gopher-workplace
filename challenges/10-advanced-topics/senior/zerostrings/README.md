# Blank Every String, However Deep

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A crash reporter serialises the request struct into the report. Legal asks for every string field to be blanked first, and the struct has four levels of nesting that change every sprint.

## Task

Implement [zerostrings.go](zerostrings.go):

1. Blank every exported string reachable from the struct `ptr` points at.
2. Descend into nested structs, pointers, interfaces, slices and arrays.
3. Leave unexported fields and non-string fields alone; a nil pointer field stays nil.
4. Return `ErrTarget` unless `ptr` is a non-nil pointer to a struct.

Replace the stub body in [zerostrings.go](zerostrings.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Redact(&record{Name:"x"})
Output: nil, Name is ""
```

**Example 2:**

```
Input:  a nested Ptr and a List of structs
Output: every Secret blanked
```

**Example 3:**

```
Input:  Redact(record{})
Output: ErrTarget
```

_Explanation:_ A value cannot be written through.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Settability propagates** | Fields reached from an addressable struct are addressable too. |
| 2 | **Recursive kind dispatch** | One case per container kind, recursing on the contents. |
| 3 | **Slice elements are addressable** | `rv.Index(i)` of a slice can be set; of an array only when the array is. |
| 4 | **CanSet as the final guard** | It covers both addressability and export status. |

## Hint

The recursive helper is written for you. Validate, step through `Elem`, and let it walk.

## Validate

```bash
make verify
```
