# Field Order Is Memory

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

The compiler never reorders struct fields — it inserts padding to satisfy alignment instead. Declare the same four fields badly and every element of a million-element slice carries eight wasted bytes, plus the cache misses that come with them.

## Task

Fill in `Record` and implement `NewRecord` in [structsizeorder.go](structsizeorder.go):

1. Declare exactly `ID int64`, `Count int32`, `Kind int16`, `Enabled bool`.
2. Order them so `unsafe.Sizeof(Record{})` is 16 bytes, not 24.
3. `NewRecord` builds a `Record` from those four values.

## Examples

**Example 1:**

```
Input:  NewRecord(1, 2, 3, true)
Output: Record{ID: 1, Count: 2, Kind: 3, Enabled: true}
```

**Example 2:**

```
Input:  unsafe.Sizeof(Record{})
Output: 16
```

**Example 3:**

```
Input:  1000 records in a slice
Output: 16000 bytes, not 24000
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Alignment drives padding** | Each field starts at a multiple of its own alignment, and the struct is padded to its widest member. |
| 2 | **Widest first** | Descending field size leaves the small fields sharing one tail word. |
| 3 | **Per-element waste multiplies** | Padding is invisible in one struct and decisive in a large slice. |

## Topics used again

Structs, `unsafe.Sizeof`, integer types.

## Hint

8, then 4, then 2, then 1 — they add up to 15 and pad to 16.

## Validate

```bash
make verify
```
