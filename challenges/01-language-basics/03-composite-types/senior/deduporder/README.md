# Dedupe Preserving Order

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

The condition is inverted: it appends only when the value was **already seen**, so
the output is exactly the duplicates.

## Task

Fix the condition between the markers in [deduporder.go](deduporder.go) to keep
first occurrences.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  [1,1,2,3,3]
Output: [1,2,3]
```

**Example 2:**

```
Input:  [5,5,5]
Output: [5]
```

**Example 3:**

```
Input:  [1,2,3]
Output: [1,2,3]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Seen set** | Track values already emitted. |
| 2 | **First occurrence** | Emit when NOT seen. |
| 3 | **Order preserved** | Slice keeps insertion order. |

## Hint

`if !ok`.

## Validate

```bash
make verify
```
