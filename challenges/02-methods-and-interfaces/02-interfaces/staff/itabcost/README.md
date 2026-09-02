# Interface Call Cost

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A hot loop dispatched through an interface. Profiling showed the calls were not inlined, and the fix required knowing when the compiler can devirtualise.

## Task

Implement the stub(s) in [itabcost.go](itabcost.go):

1. Implement `Apply` on `AddOp` and `MulOp`.
2. Implement `RunIface`, which folds a slice through an `Op` interface value.
3. Implement `RunConcrete`, which folds through a concrete `AddOp` (no dynamic dispatch).
4. Constraint: both must produce identical results and allocate zero times per call; the benchmarks expose the dispatch difference.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RunIface(AddOp{N: 2}, [1 2 3])
Output: 6+3*2 = 12
```

**Example 2:**

```
Input:  RunConcrete(AddOp{N: 2}, [1 2 3])
Output: the same 12
```

**Example 3:**

```
Input:  MulOp through the interface
Output: folded product
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **itab dispatch** | An interface call is an indirect jump through the method table — opaque to the inliner unless devirtualised. |
| 2 | **Devirtualisation** | The compiler can turn an interface call into a direct one when the dynamic type is provably fixed. |
| 3 | **Allocation-free abstraction** | Reused: dispatch is cheap; boxing values is what allocates. |

## Hint

A concrete parameter type lets the compiler inline the call; an interface parameter usually does not.

## Validate

```bash
make verify
```
