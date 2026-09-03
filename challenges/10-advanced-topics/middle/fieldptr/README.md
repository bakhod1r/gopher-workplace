# Reach A Field Through Its Offset

**Level:** middle
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A marshaller stores each field's offset once and then writes through it for every record. Getting from a struct pointer plus an offset to a typed pointer is the step that has to be right.

## Task

Implement [fieldptr.go](fieldptr.go):

1. Increment `p.Seq` by writing through a pointer built from the struct pointer and the field's offset.
2. Return the new value.
3. No other field may change.

Replace the stub body in [fieldptr.go](fieldptr.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  r := &Rec{Seq: 1}; BumpSeq(r)
Output: 2, r.Seq is 2
```

**Example 2:**

```
Input:  other fields after the call
Output: unchanged
```

**Example 3:**

```
Input:  100 calls from zero
Output: 1..100
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **unsafe.Add** | Advances an `unsafe.Pointer` by a byte count, keeping it a pointer the collector understands. |
| 2 | **unsafe.Offsetof** | The compile-time offset of the field within its struct. |
| 3 | **Pointer conversion** | Casting `unsafe.Pointer` to `*int64` is what makes the memory typed again. |

## Hint

Struct pointer to `unsafe.Pointer`, add the offset, convert to `*int64`.

## Validate

```bash
make verify
```
