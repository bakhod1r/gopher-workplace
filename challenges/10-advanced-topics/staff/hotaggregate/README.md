# Aggregate A Stream With No Garbage At All

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

An ingest hot path parses a million lines a second. Every allocation it makes is one the collector has to chase, and the profile is all allocator.

## Task

Implement [hotaggregate.go](hotaggregate.go):

1. Sum the decimal integers across every line, separated by `sep`.
2. Return the total, the field count, and `ErrSyntax` for a malformed field.
3. Empty lines are skipped; an empty field is a syntax error.
4. Zero allocations — on the success path and the error path alike.

Replace the stub body in [hotaggregate.go](hotaggregate.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Aggregate([][]byte{[]byte("1,2"), []byte("3")}, ',')
Output: 6, 3, nil
```

**Example 2:**

```
Input:  Aggregate([][]byte{[]byte("1,,2")}, ',')
Output: ErrSyntax
```

**Example 3:**

```
Input:  64 lines of 16 fields
Output: 0 allocations
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Parse in place** | Digits fold into an accumulator straight from the bytes. |
| 2 | **Sentinel errors on the failure path** | A formatted error would allocate exactly when the input is hostile. |
| 3 | **Virtual trailing separator** | Running the index to `len(line)` inclusive closes the final field. |
| 4 | **Allocation as a graded property** | `AllocsPerRun` on both paths is the specification. |

## Hint

Two nested loops and one accumulator. Nothing is ever built.

## Validate

```bash
make verify
```
