# The Format Flame Graphs Eat

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Every flame graph tool speaks the same intermediate language: one line per unique stack, frames joined by semicolons, a value at the end. Producing it is the whole bridge between a raw sample list and a picture — and it is an aggregation plus a total order, nothing more.

## Task

Implement `Fold` in [stackfold.go](stackfold.go):

1. Emit one line per distinct stack as `"frame;frame value"`, summing identical stacks.
2. Order by value descending, then by the stack string ascending.
3. Drop samples with a non-positive value or an empty stack; nothing left gives an empty, non-nil slice.

## Examples

**Example 1:**

```
Input:  [{[a b] 3} {[a b] 2}]
Output: ["a;b 5"]
```

**Example 2:**

```
Input:  [{[z] 1} {[a b] 9} {[m] 1}]
Output: ["a;b 9" "m 1" "z 1"]
```

**Example 3:**

```
Input:  [{[a b] 1} {[b a] 2}]
Output: ["b;a 2" "a;b 1"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The stack is the key** | Order matters: `a;b` and `b;a` are different call paths, not the same set. |
| 2 | **A string key for a slice** | Slices are not comparable, so the joined string is what makes the map possible. |
| 3 | **Total order for reproducibility** | Value alone leaves ties in randomised map order. |

## Topics used again

Maps, `strings.Join`, `slices.SortFunc`, `cmp.Compare`, `fmt.Sprintf`.

## Hint

Join first, aggregate into a `map[string]int64`, then sort the pairs.

## Validate

```bash
make verify
```
