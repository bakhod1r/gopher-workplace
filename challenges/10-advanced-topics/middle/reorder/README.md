# The Same Fields, Half The Bytes

**Level:** middle
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A cache holds fifty million of one small struct. Its four fields need fourteen bytes and the struct occupies twenty-four, because they were declared in the order somebody thought of them.

## Task

Fix the single planted bug in [reorder.go](reorder.go):

1. Reorder `Entry`'s fields so the struct occupies at most 16 bytes.
2. Keep every field, with its name and its type.
3. `Size` must keep reporting the type's real size.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  sizeof(Entry) as declared
Output: 24
```

_Explanation:_ Padding after each narrow field.

**Example 2:**

```
Input:  sizeof(Entry) reordered
Output: 16
```

**Example 3:**

```
Input:  the field set
Output: unchanged
```

_Explanation:_ Only the order moves.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Alignment drives padding** | Each field starts at a multiple of its alignment; the gaps are wasted. |
| 2 | **Widest first** | Descending field width usually eliminates internal padding. |
| 3 | **Tail padding** | The struct's size is rounded up to its own alignment. |

## Hint

`int64` first. Then `int32`. Then the bytes.

## Validate

```bash
make verify
```
