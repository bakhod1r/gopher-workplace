# Reach a Field by Offset

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

The offset of field B within the struct is `unsafe.Offsetof(p.B)`, not
`unsafe.Sizeof(p.B)`. Sizeof gives the field's width, not its position.

## Task

Fix [fieldoffset.go](fieldoffset.go) to use the field offset.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SecondField(&Pair{1, 2})
Output: 2
```

**Example 2:**

```
Input:  SecondField(&Pair{10, 20})
Output: 20
```

**Example 3:**

```
Input:  SecondField(&Pair{0, -5})
Output: -5
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Offsetof vs Sizeof** | Offsetof is position; Sizeof is width. |
| 2 | **Field position** | B starts at offset 4 here. |
| 3 | **unsafe.Add** | Advance base by the offset. |

## Hint

Use the field offset: `off := unsafe.Offsetof(p.B)`.

## Validate

```bash
make verify
```
