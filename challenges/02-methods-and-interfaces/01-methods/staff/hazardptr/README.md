# Hazard Pointer

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Lock-free readers have a lifetime problem: between loading a pointer and using
it, a writer may free the object. A hazard pointer is the reader's announcement
— "I am using this address" — published *before* re-checking that the pointer is
still current.

## Task

Implement `Protect` on `*Hazard` in [hazardptr.go](hazardptr.go):

1. Load the pointer from `shared`.
2. Store it into `h.ptr` — the hazard announcement.
3. Load `shared` again; if it still holds the same pointer, return it.
4. Otherwise return `nil` — the caller must retry.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  shared holds &val (42); Protect(&shared)
Output: the same pointer; *got == 42
```

**Example 2:**

```
Input:  shared changed between the two loads
Output: nil  (caller retries)
```

**Example 3:**

```
Input:  shared holds nil
Output: nil
```

_Explanation:_ the validation compares the two loads; nothing else can be assumed.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`atomic.Pointer[T]`** | Typed atomic load/store — no `unsafe.Pointer` casts needed. |
| 2 | **Announce, then validate** | Publishing the hazard *before* the second load is the entire protocol. |
| 3 | **Retry as the failure mode** | A lock-free reader never blocks; it re-reads. |

## Hint

The order is load → store → load → compare. Storing after the validation would
leave a window in which the reclaimer cannot see your announcement.

## Validate

```bash
make verify
```
