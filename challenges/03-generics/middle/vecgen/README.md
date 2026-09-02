# Vector Arithmetic

**Level:** middle  
**Topic:** 03-generics

## Context

A recommendation score is a dot product over feature vectors that are ints in one pipeline and floats in another.

## Task

Implement the stub(s) in [vecgen.go](vecgen.go):

1. Implement `Add` and `Dot`.
2. Mismatched lengths yield an empty slice and `false` respectively — never a panic.
3. Both work for any numeric element type.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Add([]int{1,2}, []int{3,4})
Output: []int{4,6}
```

**Example 2:**

```
Input:  Dot([]int{1,2}, []int{3,4})
Output: 11, true
```

**Example 3:**

```
Input:  Dot([]int{1}, []int{1,2})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Numeric constraints** | `+` and `*` both come from the type set. |
| 2 | **Shape validation** | Length equality is the vector-space precondition; check it once, up front. |
| 3 | **Accumulating in `T`** | The dot product keeps the element type, so no conversion is needed. |

## Hint

Validate the lengths first, then the loops need no bounds thinking.

## Validate

```bash
make verify
```
