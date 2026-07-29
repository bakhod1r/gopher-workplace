# Copy Bounded by Destination

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

`copy(dst, src)` copies `min(len(dst), len(src))` elements. A destination of
length 0 (even with capacity) copies nothing; the length, not the capacity, is
what counts.

## Task

Fix [copyshort.go](copyshort.go) so the clone actually contains the data.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Clone([1 2 3])
Output: [1 2 3], independent
```

**Example 2:**

```
Input:  len(Clone([1 2 3]))
Output: 3
```

**Example 3:**

```
Input:  Clone(nil)
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **copy is length-bounded** | It copies min(len(dst), len(src)). |
| 2 | **length vs capacity** | Capacity alone does not make room for copy. |
| 3 | **make length** | Allocate `len(xs)` elements. |

## Hint

Make the destination with LENGTH len(xs): `dst := make([]int, len(xs))`.

## Validate

```bash
make verify
```
