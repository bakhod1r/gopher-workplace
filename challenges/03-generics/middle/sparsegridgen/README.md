# Sparse Grid

**Level:** middle  
**Topic:** 03-generics

## Context

A map editor covers a huge coordinate space of which almost nothing is painted, so a dense matrix would waste most of its memory.

## Task

Implement the stub(s) in [sparsegridgen.go](sparsegridgen.go):

1. Implement `NewGrid`, `Set`, `At`, and `Filled`.
2. Unset cells read as the default, including negative coordinates.
3. A cell explicitly set to the default still counts as filled.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Set(1,1,x); At(1,1)
Output: x
```

**Example 2:**

```
Input:  At(9,9) unset
Output: the default
```

**Example 3:**

```
Input:  Set(0,0,def); Filled()
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Struct map keys** | A comparable struct like `point` makes a natural composite key. |
| 2 | **Sparse versus dense** | Memory grows with the number of set cells, not the coordinate range. |
| 3 | **Unbounded coordinates** | Negative and huge coordinates work without any allocation. |

## Hint

A small comparable struct is the key — no string formatting needed.

## Validate

```bash
make verify
```
