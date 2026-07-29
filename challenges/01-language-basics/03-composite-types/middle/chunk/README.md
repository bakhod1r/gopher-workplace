# Chunk a Slice

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

A batch processor sends records in groups of N. Split a slice into fixed-size
chunks, the last possibly short.

## Task

Implement `Chunk(xs, size)`; `size <= 0` → empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,2,3,4,5], size=2
Output: [[1,2],[3,4],[5]]
```

_Explanation:_ last chunk shorter

**Example 2:**

```
Input:  [1,2,3], size=5
Output: [[1,2,3]]
```

_Explanation:_ size exceeds len

**Example 3:**

```
Input:  [1], size=0
Output: []
```

_Explanation:_ size<=0 gives empty

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Slice windows** | `xs[i:end]` where end is clamped. |
| 2 | **Step by size** | Advance i by size each iteration. |
| 3 | **Clamp last** | `min(i+size, len(xs))`. |

## Hint

`for i := 0; i < len(xs); i += size { end := i+size; if end > len(xs) { end = len(xs) }; out = append(out, xs[i:end]) }`.

## Validate

```bash
make verify
```
