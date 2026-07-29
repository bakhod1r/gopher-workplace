# Reinterpret Struct as Array

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

Reinterpreting `*Pair` as `*[2]int32` gives an array view whose elements are the
two fields (same layout). The bug returns only `arr[0]`; sum both slots.

## Task

Fix [structasarray.go](structasarray.go) to sum both reinterpreted elements.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Sum(&Pair{3, 4})
Output: 7
```

**Example 2:**

```
Input:  Sum(&Pair{0, 0})
Output: 0
```

**Example 3:**

```
Input:  Sum(&Pair{-1, 5})
Output: 4
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Struct/array layout equivalence** | Two int32 fields lay out like [2]int32. |
| 2 | **Reinterpret view** | `(*[2]int32)(unsafe.Pointer(p))`. |
| 3 | **Read all elements** | arr[0] + arr[1]. |

## Hint

Sum both: `return arr[0] + arr[1]`.

## Validate

```bash
make verify
```
