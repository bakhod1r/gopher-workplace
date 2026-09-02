# Min Max Tracker

**Level:** junior  
**Topic:** 03-generics

## Context

A streaming gauge sees values one at a time and can never hold the whole series in memory.

## Task

Implement the stub(s) in [trackergen.go](trackergen.go):

1. Implement `Add`, updating the recorded range.
2. Implement `Bounds`, returning the smallest and largest values and `true`.
3. Before anything is added, `Bounds` returns zero values and `false`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Add(3); Add(1); Bounds()
Output: 1, 3, true
```

**Example 2:**

```
Input:  Add(5); Bounds()
Output: 5, 5, true
```

**Example 3:**

```
Input:  Bounds() before any Add
Output: 0, 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Constrained type parameters on types** | `Tracker[T cmp.Ordered]` restricts what the type can be instantiated with. |
| 2 | **Seeding on the first value** | The first `Add` must set both bounds, not compare against zero values. |
| 3 | **Pointer receivers mutate** | Reused from methods: a value receiver gets a copy, so mutation needs `*T`. |

## Hint

The first `Add` is a special case: it seeds both bounds.

## Validate

```bash
make verify
```
