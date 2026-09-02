# Nil Versus Empty

**Level:** junior  
**Topic:** 03-generics

## Context

A JSON endpoint rendered `null` instead of `[]` whenever a filter matched nothing, breaking the frontend.

## Task

Implement the stub(s) in [nilslicepitfall.go](nilslicepitfall.go):

1. Implement `Collect` so the result is never nil, even when nothing matches.
2. Implement `IsNil`, reporting whether a slice is nil.
3. Note that `len` cannot tell nil from empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Collect([]int{1}, none)
Output: []int{} (not nil)
```

**Example 2:**

```
Input:  IsNil(Collect(...))
Output: false
```

**Example 3:**

```
Input:  IsNil([]int(nil))
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nil and empty differ** | Both have length 0, but they marshal differently and compare differently to nil. |
| 2 | **`var out []T` versus `make`** | `var` gives a nil slice; `make(...)` gives an empty, non-nil one. |
| 3 | **Generic code inherits this** | The rule is not special to type parameters — but generic helpers hide it from callers. |

## Hint

Start with `make([]T, 0, len(s))`, not `var out []T`.

## Validate

```bash
make verify
```
