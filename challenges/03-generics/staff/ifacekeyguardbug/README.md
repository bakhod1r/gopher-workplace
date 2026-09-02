# The Set That Panics On A Slice

**Level:** staff  
**Topic:** 03-generics

## Context

A deduplication helper instantiated at `any` crashes the request handler with `hash of unhashable type []int` whenever a caller passes a decoded JSON array through it.

## Task

Fix the single planted bug in [ifacekeyguardbug.go](ifacekeyguardbug.go):

1. Find and fix the single bug so an element with an uncomparable dynamic type is passed through instead of panicking.
2. Deduplication of ordinary values, and input order, must not change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Distinct([]any{1, 1, 2})
Output: []any{1, 2}
```

**Example 2:**

```
Input:  Distinct([]any{1, []int{2}, 1})
Output: []any{1, []int{2}}
```

**Example 3:**

```
Input:  Distinct([]string{"a", "a"})
Output: []string{"a"}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **comparable is not strictly comparable** | Since Go 1.20 an ordinary interface type *satisfies* `comparable` even though comparing it can panic at run time. |
| 2 | **Dynamic versus static** | The compiler checks the static type parameter; the hash is computed on the dynamic type inside the interface. |
| 3 | **Guard, don't hope** | Any generic map keyed on a type parameter that may be instantiated at an interface needs a run-time guard. |

## Hint

There is a helper in this file that nothing calls.

## Validate

```bash
make verify
```
