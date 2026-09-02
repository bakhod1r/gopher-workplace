# Allocate Once

**Level:** middle  
**Topic:** 03-generics

## Context

A request handler builds a response slice of known size. Growing it by repeated appends showed up as garbage-collector pressure in profiles.

## Task

Implement the stub(s) in [preallocgen.go](preallocgen.go):

1. Implement `Build`, producing `n` elements from `f(i)`.
2. Allocate exactly once: after the call the capacity must equal `n`.
3. Treat a negative `n` as zero.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Build(3, f)
Output: 3 elements, cap 3
```

**Example 2:**

```
Input:  Build(0, f)
Output: empty, cap 0
```

**Example 3:**

```
Input:  Build(-1, f)
Output: empty
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Capacity hints** | `make([]T, 0, n)` reserves once; the appends never reallocate. |
| 2 | **Length versus capacity** | Starting at length 0 with capacity `n` keeps `append` writing into reserved space. |
| 3 | **Measurable, not theoretical** | The test asserts the resulting capacity, which is what the reservation buys. |

## Hint

`make([]T, 0, n)` — length zero, capacity `n`.

## Validate

```bash
make verify
```
