# Check The Alignment You Depend On

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A struct is laid out so its counter can be updated atomically. The invariant is asserted with a literal 8, and on a 32-bit build the assertion is both wrong and green.

## Task

Fix the single planted bug in [alignassert.go](alignassert.go):

1. Report whether `Counter.Value` is aligned for atomic access.
2. Derive the requirement from `unsafe.Alignof`, not from a literal.
3. Check the field's offset within the struct and the address of an actual instance.
4. Fix the single bug.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Check()
Output: true
```

_Explanation:_ Value is the first field of a well-aligned struct.

**Example 2:**

```
Input:  100 runs
Output: true every time
```

**Example 3:**

```
Input:  the requirement's source
Output: unsafe.Alignof
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **unsafe.Alignof** | The type's required alignment, as a compile-time constant. |
| 2 | **Alignment is per platform** | A 64-bit value is not 8-byte aligned on every architecture by default. |
| 3 | **Offset and address** | Both matter: a well-placed field in a misaligned struct is still misaligned. |

## Hint

The number 8 is an answer, not a question. Which call asks the question?

## Validate

```bash
make verify
```
