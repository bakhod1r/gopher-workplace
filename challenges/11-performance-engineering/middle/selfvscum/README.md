# Both Columns, One Pass

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Reading a profile means holding two numbers per function at once: flat says who burns CPU, cum says who is responsible for it. A function with high cum and near-zero flat is a coordinator — optimise its callees. High flat is where the instructions actually retire.

## Task

Implement both functions in [selfvscum.go](selfvscum.go):

1. `Analyze` credits `Flat` to the leaf frame only and `Cum` to every *distinct* frame on the stack, in one pass.
2. Ignore samples with a non-positive value or an empty stack.
3. `Leaves` returns the functions with `Flat >= minFlat`, ordered by flat descending then name ascending.

## Examples

**Example 1:**

```
Input:  Analyze([{[main a] 5}])
Output: {main:{0 5} a:{5 5}}
```

**Example 2:**

```
Input:  Analyze([{[rec rec rec] 6}])
Output: {rec:{6 6}}
```

**Example 3:**

```
Input:  Leaves({a:{5 5} b:{0 9} c:{5 5} d:{1 1}}, 1)
Output: [a c d]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Flat and cum answer different questions** | "Who is slow" versus "what is this time spent on". |
| 2 | **Recursion inflates cum** | Without a per-sample seen-set, a deep recursion multiplies its own cum. |
| 3 | **One pass, two accumulators** | Reading the sample list twice is a habit worth breaking on large profiles. |

## Topics used again

Maps of structs, sets, `slices.SortFunc`, `cmp.Compare`.

## Hint

Updating a struct in a map means read, modify, write back — you cannot assign to `m[k].Flat`.

## Validate

```bash
make verify
```
