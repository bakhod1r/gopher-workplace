# Interface Slices

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A gateway receives concrete records and must hand them to code that only accepts the interface.

## Task

Implement the stub(s) in [ifaceslice.go](ifaceslice.go):

1. Implement `ID` on `User` and `Order`.
2. Implement `ToEntities`, which converts a `[]User` into a `[]Entity` — Go will not convert slices implicitly.
3. Implement `IDs`, which returns the ids of a mixed `[]Entity`.
4. Preallocate the result slice; do not grow it one element at a time from zero capacity.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ToEntities([]User{{ID_: "u1"}})
Output: []Entity of length 1
```

**Example 2:**

```
Input:  IDs([]Entity{User{ID_: "u1"}, Order{ID_: "o1"}})
Output: ["u1", "o1"]
```

**Example 3:**

```
Input:  ToEntities(nil)
Output: empty slice
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **No implicit slice conversion** | `[]User` is not assignable to `[]Entity`, even when `User` satisfies `Entity`. |
| 2 | **Memory layout** | An interface element is a two-word header, so the slices have different shapes. |
| 3 | **Preallocation** | Efficiency: `make` with the known length avoids repeated regrowth. |

## Hint

Convert element by element — the slice types have different memory layouts, so a cast is impossible.

## Validate

```bash
make verify
```
