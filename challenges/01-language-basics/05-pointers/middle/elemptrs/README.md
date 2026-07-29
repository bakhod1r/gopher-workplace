# Pointers Into a Slice

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

Taking `&xs[i]` yields a pointer that aliases the slice element; mutating
through it changes the slice. (Go 1.22 per-iteration range vars make this safe.)

## Task

Implement `Pointers` in [elemptrs.go](elemptrs.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  xs := []int{1,2,3}; ps := Pointers(xs); *ps[1] = 99
Output: xs[1] == 99
```

_Explanation:_ Each pointer aliases a real slice element.

**Example 2:**

```
Input:  Pointers([]int{})
Output: [] (empty)
```

**Example 3:**

```
Input:  ps := Pointers(xs); *ps[0]
Output: xs[0]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Address of element** | `&xs[i]` aliases the slice slot. |
| 2 | **Slice of pointers** | Collect each address. |
| 3 | **Aliasing** | Writes propagate to xs. |

## Hint

Range with index: `for i := range xs { ps = append(ps, &xs[i]) }`.

## Validate

```bash
make verify
```
