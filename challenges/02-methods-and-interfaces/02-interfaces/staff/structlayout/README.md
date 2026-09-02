# Struct Layout

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A 100M-element slice of records used far more memory than the field sizes suggested. Field ordering was the difference.

## Task

Implement the stub(s) in [structlayout.go](structlayout.go):

1. Define `Packed` with the same fields as `Padded` but ordered so the struct is smaller.
2. Implement `Size` on both, and `TotalBytes`, which reports the memory a slice of n elements occupies.
3. Constraint: `unsafe.Sizeof(Packed{})` must be strictly smaller than `unsafe.Sizeof(Padded{})`, which the test asserts.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  unsafe.Sizeof(Padded{})
Output: 24 on a 64-bit build
```

**Example 2:**

```
Input:  unsafe.Sizeof(Packed{})
Output: 16
```

**Example 3:**

```
Input:  TotalBytes for 1M packed records
Output: a third less than padded
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Alignment and padding** | Each field is aligned to its own size; gaps are inserted to satisfy that. |
| 2 | **Field ordering** | Descending size order minimises padding. |
| 3 | **unsafe.Sizeof / Offsetof** | Reused: layout claims verified mechanically. |

## Hint

Order fields largest-alignment first: pointers and 8-byte values, then 4, then 2, then 1.

## Validate

```bash
make verify
```
