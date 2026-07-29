# Array Pointer to Slice

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

Slicing an array pointer `p[:]` yields a slice over ALL elements, aliasing the
array. `p[:2]` drops the last two.

## Task

Fix [arrtoslice.go](arrtoslice.go) to view the whole array.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  AsSlice(&[4]int{1,2,3,4})
Output: len 4, aliases the array
```

**Example 2:**

```
Input:  len(AsSlice(&[4]int{}))
Output: 4
```

**Example 3:**

```
Input:  AsSlice(&a)[0] = 9
Output: a[0] == 9
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Array pointer slicing** | `p[:]` covers the whole array. |
| 2 | **Aliasing** | The slice shares the array's storage. |
| 3 | **Full length** | Use `p[:]`, not a fixed short length. |

## Hint

Slice the whole array: `return p[:]`.

## Validate

```bash
make verify
```
