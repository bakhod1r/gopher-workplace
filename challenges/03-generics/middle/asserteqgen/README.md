# Slice Assertion Helper

**Level:** middle  
**Topic:** 03-generics

## Context

Slice comparisons in tests either use `reflect.DeepEqual` and print two walls of text, or get written by hand in every file.

## Task

Implement the stub(s) in [asserteqgen.go](asserteqgen.go):

1. Implement `EqualSlice`, reporting a precise failure.
2. Report a length mismatch separately from an element mismatch, naming the index.
3. Return whether the slices matched.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  EqualSlice(t, []int{1}, []int{1})
Output: true, no failure
```

**Example 2:**

```
Input:  different lengths
Output: false, one failure mentioning the lengths
```

**Example 3:**

```
Input:  differing element
Output: false, one failure naming the index
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Precise failures** | Naming the index turns a diff into a pointer at the problem. |
| 2 | **`testing.TB`** | Accepting the interface lets benchmarks use the helper too. |
| 3 | **Typed assertions** | `comparable` catches mismatched element types at compile time. |

## Hint

Report once and return — do not flood the output with every difference.

## Validate

```bash
make verify
```
