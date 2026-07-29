# Struct With Pointer Field Size

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

A pointer is 8 bytes; a struct of two pointers is 16. Measuring one pointer
returns 8. Measure the struct: `unsafe.Sizeof(Pair{})`.

## Task

Fix [ptrfieldsize.go](ptrfieldsize.go) to return the struct size.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Size()
Output: 16
```

_Explanation:_ A struct of two 8-byte pointers.

**Example 2:**

```
Input:  Sizeof(Pair{})
Output: 16
```

**Example 3:**

```
Input:  Sizeof((*int)(nil))
Output: 8 (wrong answer)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer size** | 8 bytes on 64-bit. |
| 2 | **Struct of pointers** | Sum of the pointer fields. |
| 3 | **Measure the struct** | `Sizeof(Pair{})`. |

## Hint

Measure the struct: `return unsafe.Sizeof(Pair{})`.

## Validate

```bash
make verify
```
