# The Accumulator That Escaped Through A Callback

**Level:** senior
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A tidy refactor replaces a loop with a callback helper. The behaviour is identical, the benchmark is 4x slower, and the allocation count went from zero to two.

## Task

Fix the single planted bug in [escapecallback.go](escapecallback.go):

1. Return the sum of `s` as an int64.
2. Fix the single bug so the function allocates nothing.
3. `Each` must stay as it is — it is a fixed part of the package.
4. Do not change `Each` into a plain function to dodge the problem.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Sum([]int{1,2,3})
Output: 6
```

**Example 2:**

```
Input:  Sum(nil)
Output: 0
```

**Example 3:**

```
Input:  512 elements
Output: 0 allocations
```

_Explanation:_ The accumulator must stay in the frame.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Closures escape into non-inlined calls** | The callee may store the func value, so the compiler must assume it does. |
| 2 | **Capture drags the variable along** | `total` is captured by reference, so it escapes with the closure. |
| 3 | **Inlining is the enabling optimisation** | An inlined callback often costs nothing; a `//go:noinline` one never does. |

## Hint

`Each` is a func variable, not a func. What does that force the compiler to assume about the closure it is given?

## Validate

```bash
make verify
```
