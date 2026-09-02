# Sequence Lock

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A telemetry snapshot is read far more often than it is written. Readers must never block writers and must be able to detect a torn read.

## Task

Implement the stub(s) in [seqlockifc.go](seqlockifc.go):

1. Implement `Write` on `*SeqLock`, bumping the sequence to odd before the update and to even after it.
2. Implement `Read`, which retries until it observes a stable even sequence around a consistent snapshot.
3. Constraint: `-race` clean, readers never block, and a returned snapshot must never mix fields from two different writes.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Write then Read
Output: the written values
```

**Example 2:**

```
Input:  reads concurrent with writes
Output: every snapshot is self-consistent
```

**Example 3:**

```
Input:  Read before any Write
Output: the zero snapshot
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Seqlock** | Optimistic reads validated by a version counter. |
| 2 | **Odd/even versioning** | An odd sequence marks a write in progress. |
| 3 | **Atomics for ordering** | Reused: atomic loads and stores establish the visibility the algorithm relies on. |

## Hint

Read the sequence, read the fields, read the sequence again — retry unless both reads are equal and even.

## Validate

```bash
make verify
```
