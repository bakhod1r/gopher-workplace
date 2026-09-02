# The Cost Of any

**Level:** middle  
**Topic:** 03-generics

## Context

An older helper took `[]any` and asserted each element. The generic replacement keeps types static — and the old one shows why that matters.

## Task

Implement the stub(s) in [boxingcostgen.go](boxingcostgen.go):

1. Implement `SumTyped` for a typed slice.
2. Implement `SumAny`, which must report `false` on the first element that is not an `int`.
3. Compare the two: one cannot fail at run time, the other can.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SumTyped([]int{1,2})
Output: 3
```

**Example 2:**

```
Input:  SumAny([]any{1,2})
Output: 3, true
```

**Example 3:**

```
Input:  SumAny([]any{1,"x"})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Static versus dynamic** | The generic version cannot be handed a wrong element type; the `any` version can. |
| 2 | **Boxing** | Every element in `[]any` carries a type word — more memory and an indirection per read. |
| 3 | **Failure moves to compile time** | The generic version turns a run-time check into a compile error. |

## Hint

The generic version has no error path — that is the point of the comparison.

## Validate

```bash
make verify
```
