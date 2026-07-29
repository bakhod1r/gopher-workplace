# Reinterpret Float Bits

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

The bug narrows to float32 first (losing precision) and reads 32 bits. Read the
full float64 pattern: reinterpret `&f` as `*uint64`.

## Task

Fix [floatbits.go](floatbits.go) to return the bit pattern.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Bits(1.5)
Output: 0x3ff8000000000000
```

**Example 2:**

```
Input:  Bits(0.0)
Output: 0x0
```

**Example 3:**

```
Input:  Bits(-2.0)
Output: 0xc000000000000000
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Value cast vs reinterpret** | `uint64(f)` truncates the number. |
| 2 | **unsafe.Pointer reinterpret** | `*(*uint64)(unsafe.Pointer(&f))`. |
| 3 | **Same-size types** | float64 and uint64 are both 8 bytes. |

## Hint

Reinterpret the bits: `return *(*uint64)(unsafe.Pointer(&f))`.

## Validate

```bash
make verify
```
