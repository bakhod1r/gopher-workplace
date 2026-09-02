# Lazy Initialization

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Some values are expensive and often unused. `LazyString` defers the work until
the first read, using a nil pointer as the "not computed yet" marker — which,
unlike a zero string, cannot be confused with a legitimate result.

## Task

Implement `String` on `*LazyString` in [lazyinit.go](lazyinit.go):

1. If `l.val` is nil, call `l.init()` and store a pointer to the result.
2. Return the dereferenced value.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  New(init); no call yet
Output: init has not run (call count 0)
```

**Example 2:**

```
Input:  first String()
Output: "heavy"  (init called once)
```

**Example 3:**

```
Input:  second String()
Output: "heavy"  (init not called again)
```

_Explanation:_ once `val` is non-nil it is the answer forever.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nil pointer as a sentinel** | `*string` distinguishes "unset" from "set to empty string". |
| 2 | **`String()` method** | This signature also satisfies `fmt.Stringer`, so the type prints lazily too. |
| 3 | **Taking the address of a local** | `v := l.init(); l.val = &v` — Go moves `v` to the heap automatically. |

## Hint

You cannot write `l.val = &l.init()`; you cannot take the address of a call
result. Assign to a local first, then take its address.

## Validate

```bash
make verify
```
