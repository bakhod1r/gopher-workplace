# Size And Alignment Of A Run-Time Type

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A layout auditor wants a table of every registered type's footprint. `unsafe.Sizeof` needs the type at compile time, and the registry only has values.

## Task

Implement [typesizes.go](typesizes.go):

1. Return the size and alignment of `v`'s dynamic type.
2. Report false for a nil interface.
3. Measure the type itself — a slice's size is its header, not its elements.

Replace the stub body in [typesizes.go](typesizes.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Sizes(int64(0))
Output: 8, 8, true
```

**Example 2:**

```
Input:  Sizes([]int{1,2,3})
Output: 24, 8, true
```

_Explanation:_ The header, not the three ints.

**Example 3:**

```
Input:  Sizes(nil)
Output: 0, 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type.Size and Type.Align** | The run-time equivalents of `unsafe.Sizeof` and `unsafe.Alignof`. |
| 2 | **Align returns an int** | Unlike `Size`, so it needs a conversion. |
| 3 | **Headers vs payload** | Reflection measures the same thing the compiler would. |

## Hint

`Size` is already a uintptr. `Align` is not.

## Validate

```bash
make verify
```
