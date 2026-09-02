# Atomic Bitset

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A presence set is updated by many goroutines. A mutex around it was the hot spot; the bits are now updated with CAS on 64-bit words.

## Task

Implement the stub(s) in [atomicbitset.go](atomicbitset.go):

1. Implement `Set`, `Clear`, `Test`, and `Count` on `*Bitset` using atomic operations on `uint64` words.
2. `Set` returns whether the bit changed from 0 to 1, `Clear` whether it changed from 1 to 0.
3. Constraint: `-race` clean, no mutex, and concurrent `Set` calls on distinct bits in the same word must not lose updates.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Set(5) twice
Output: true, then false
```

**Example 2:**

```
Input:  Test(5)
Output: true
```

**Example 3:**

```
Input:  64 goroutines setting bits 0..63
Output: Count 64
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **CAS on words** | A bit update is a read-modify-write of the whole word. |
| 2 | **Lost updates** | A plain `word |= mask` on a shared word drops concurrent bits. |
| 3 | **popcount** | Reused: counting set bits with `math/bits`. |

## Hint

`for { old := w.Load(); new := old | mask; if old == new { return false }; if w.CompareAndSwap(old, new) { return true } }`.

## Validate

```bash
make verify
```
