# The Frames That Are Not Really There

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Inlining is the optimisation that makes small Go functions free — and the reason profiles can look wrong. An inlined function has no stack frame; the profiler reconstructs its name from the inline tree the compiler recorded. That reconstruction is what lets you see `slices.Contains` in a profile at all, and hiding it again is what `pprof -noinlines` does when you want to know which machine code is actually hot.

## Task

Implement both functions in [inlinemark.go](inlinemark.go):

1. `Physical` returns the names of the non-inlined frames, in order, without modifying the input.
2. An all-inlined or empty stack gives an empty, non-nil slice.
3. `Attribute` returns the innermost non-inlined frame — the function that owns the leaf's machine code — and reports `"", false` when there is none.

## Examples

**Example 1:**

```
Input:  Physical([{a false} {b true} {c false}])
Output: [a c]
```

**Example 2:**

```
Input:  Attribute([{a false} {b true}])
Output: "a", true
```

**Example 3:**

```
Input:  Attribute([{a true} {b true}])
Output: "", false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Inlined frames are metadata** | They exist in the inline tree, not on the stack the CPU walked. |
| 2 | **Attribution follows the machine code** | The samples landed in the caller's body, whatever names the tree reports. |
| 3 | **Inlining explains a "missing" function** | A hot helper that never appears was folded into every caller. |

## Topics used again

Structs, slices, reverse iteration, multiple return values.

## Hint

`Attribute` scans from the leaf end backwards and stops at the first non-inlined frame.

## Validate

```bash
make verify
```
