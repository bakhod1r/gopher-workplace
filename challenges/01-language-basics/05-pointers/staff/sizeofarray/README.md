# Size of Whole Array

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

`unsafe.Sizeof(p[0])` is one element (4 bytes). The whole array is
`unsafe.Sizeof(*p)` = 4 * 4 = 16 bytes.

## Task

Fix [sizeofarray.go](sizeofarray.go) to return the whole array's size.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TotalSize(&[4]int32{})
Output: 16
```

**Example 2:**

```
Input:  TotalSize(&[2]int64{})
Output: 16
```

**Example 3:**

```
Input:  TotalSize(&[1]byte{})
Output: 1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Array size** | `Sizeof(*p)` covers all elements. |
| 2 | **Element vs array** | `p[0]` is one element. |
| 3 | **Fixed length in the type** | Array length is part of its type. |

## Hint

Measure the whole array: `return unsafe.Sizeof(*p)`.

## Validate

```bash
make verify
```
