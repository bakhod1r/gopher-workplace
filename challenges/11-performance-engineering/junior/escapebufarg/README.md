# Returning A Value Costs Nothing

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Every allocation starts as an escape-analysis decision: can this value stay on the stack, or must the heap hold it? A small struct returned by value stays on the stack. The same struct returned as `*Point` escapes, and now the garbage collector has one more object per call.

## Task

Implement both functions in [escapebufarg.go](escapebufarg.go):

1. `Add` returns the component-wise sum by value, without touching its inputs.
2. `AddInto` writes the same sum through `dst`.
3. Neither may allocate, including when `dst` also appears as an input.

## Examples

**Example 1:**

```
Input:  Add({1 2}, {3 4})
Output: {4 6}
```

**Example 2:**

```
Input:  AddInto(&p, {1 2}, {3 4})
Output: p is {4 6}
```

**Example 3:**

```
Input:  p is {4 6}; AddInto(&p, p, {1 1})
Output: p is {5 7}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Escape analysis** | A value whose address outlives the call must live on the heap. |
| 2 | **Small structs are cheap to copy** | Two words move in registers; a pointer costs an allocation plus an indirection. |
| 3 | **Write-through APIs** | Passing the destination lets the caller decide where the memory lives. |

## Topics used again

Structs, pointers, value semantics.

## Hint

Both functions are one line each; neither needs `new` or `&Point{}`.

## Validate

```bash
make verify
```
