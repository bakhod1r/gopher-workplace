# Cloning A Matrix

**Level:** middle  
**Topic:** 03-generics

## Context

An undo feature snapshots a grid before each edit. Copying only the outer slice made every snapshot show the latest state.

## Task

Implement the stub(s) in [clone2dgen.go](clone2dgen.go):

1. Implement `Clone2D`, copying the outer slice and every row.
2. Writing into a copied row must not affect the original.
3. Return an empty (non-nil) result for empty or nil input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Clone2D([][]int{{1,2}})
Output: [][]int{{1,2}}
```

**Example 2:**

```
Input:  writing into a copied row
Output: original unchanged
```

**Example 3:**

```
Input:  Clone2D(nil)
Output: [][]int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nested copies** | `slices.Clone` on `[][]T` copies the row headers, not the rows. |
| 2 | **Row independence** | Each row needs its own allocation for the snapshot to be meaningful. |
| 3 | **Empty rows** | A zero-length row still gets its own (empty) slice. |

## Hint

One `make`+`copy` per row, not one for the whole thing.

## Validate

```bash
make verify
```
