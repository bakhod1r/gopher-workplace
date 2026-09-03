# Call A Function You Only Know At Run Time

**Level:** middle
**Topic:** 10-advanced-topics / 03-reflection

## Context

A plugin registry stores handler functions as `any`. Calling one means checking, at run time, that it has the shape the registry promised.

## Task

Implement [callfn.go](callfn.go):

1. Call `fn` with `args` and return its int results.
2. Reject anything that is not a function, has the wrong arity, is variadic, or has a non-int parameter or result.
3. A function with no results returns an empty slice, not an error.

Replace the stub body in [callfn.go](callfn.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  CallInts(func(a, b int) int { return a+b }, 1, 2)
Output: [3], nil
```

**Example 2:**

```
Input:  CallInts(func(a int) (int,int) { return a,-a }, 5)
Output: [5 -5], nil
```

_Explanation:_ Every result is collected.

**Example 3:**

```
Input:  CallInts(func(s string) int {...}, 1)
Output: ErrSignature
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **reflect.Value.Call** | Takes and returns `[]reflect.Value`; a mismatch panics, so validate first. |
| 2 | **Type.NumIn / In / NumOut / Out** | The function's signature is fully inspectable. |
| 3 | **IsVariadic** | A variadic function needs `CallSlice`, not `Call` — reject it here. |
| 4 | **Turning panics into errors** | Validating up front is better than recovering afterwards. |

## Hint

Every way `Call` can panic corresponds to a check you can make on the `Type` first.

## Validate

```bash
make verify
```
