# Keys That Were Overwritten By The Next Batch

**Level:** staff
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A column store slices its read arena into keys and inserts them into a map. Lookups start missing entries that are visibly present, and the map's length grows past the number of distinct keys.

## Task

Fix the single planted bug in [arenastrings.go](arenastrings.go):

1. Return one string per span, for the caller to keep past the next batch.
2. An out-of-range, inverted or empty span yields the empty string.
3. Fix the single bug so the strings do not view the arena.
4. Copy the batch into one block: a handful of allocations for 32 spans, not one per span.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Intern([]byte("abcd"), [][2]int{{0,2},{2,4}})
Output: ["ab" "cd"]
```

**Example 2:**

```
Input:  keys := Intern(arena, ...); copy(arena, "OVERWRIT")
Output: keys unchanged
```

**Example 3:**

```
Input:  a span of {3,1}
Output: ""
```

_Explanation:_ Inverted spans are rejected.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Map keys must be stable** | A map caches each key's hash; mutating a key's bytes strands the entry. |
| 2 | **Arena copying** | One block plus per-span headers beats one allocation per span. |
| 3 | **unsafe.String over your own block** | Legal precisely because nothing else can write to it. |
| 4 | **Sub-slicing the block** | Each string wraps `block[start:]` with its own length. |

## Hint

The conversion is right; the bytes are the caller's. Where should the batch's bytes live?

## Validate

```bash
make verify
```
