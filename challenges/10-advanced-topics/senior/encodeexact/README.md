# Serialise With The Size You Can Compute

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

An export path builds its payload by appending to a nil slice. For a large batch that is a dozen reallocations and a full copy at each one.

## Task

Implement [encodeexact.go](encodeexact.go):

1. Render each record as `id:name`, separated by `\n`.
2. Compute the exact output length first and allocate once.
3. An empty input returns an empty result.

Replace the stub body in [encodeexact.go](encodeexact.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Encode([]Rec{{1,"a"},{22,"bb"}})
Output: "1:a\n22:bb"
```

**Example 2:**

```
Input:  Encode([]Rec{{-5,""}})
Output: "-5:"
```

_Explanation:_ The sign counts; an empty name is legal.

**Example 3:**

```
Input:  64 records
Output: 1 allocation
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two passes beat doubling** | One cheap sizing pass removes every reallocation. |
| 2 | **Digit counting** | `len(strconv.Itoa(id))` is the simplest correct width, sign included. |
| 3 | **Separator arithmetic** | n records need n-1 separators. |
| 4 | **strconv.Append*** | Writes digits straight into the buffer. |

## Hint

The separators are part of the length too.

## Validate

```bash
make verify
```
