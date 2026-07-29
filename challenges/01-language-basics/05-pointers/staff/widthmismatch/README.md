# Reinterpret With Wrong Width

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

Reinterpreting an 8-byte int64 through `*uint32` reads only 4 bytes. The target
type must match the source width: reinterpret through `*uint64`.

## Task

Fix [widthmismatch.go](widthmismatch.go) to read all 64 bits.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  AsU64(-1)
Output: 0xffffffffffffffff
```

**Example 2:**

```
Input:  AsU64(0)
Output: 0x0
```

**Example 3:**

```
Input:  AsU64(1)
Output: 0x1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Matching widths** | Source and target must be the same size. |
| 2 | **Truncated read** | *uint32 reads 4 of 8 bytes. |
| 3 | **Correct target type** | *uint64. |

## Hint

Reinterpret through the same-width type: `return *(*uint64)(unsafe.Pointer(&x))`.

## Validate

```bash
make verify
```
