# Max Of Slice

**Level:** junior  
**Topic:** 03-generics

## Context

A monitoring panel reports the peak value of a sample window, which is sometimes empty.

## Task

Implement the stub(s) in [maxslice.go](maxslice.go):

1. Implement `MaxOf`, returning the largest element and `true`.
2. Return the zero value and `false` for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MaxOf([]int{1, 9, 3})
Output: 9, true
```

**Example 2:**

```
Input:  MaxOf([]string{"a", "c"})
Output: "c", true
```

**Example 3:**

```
Input:  MaxOf([]int{})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<`, `<=`, `>`, `>=`. |
| 2 | **Zero value of `T`** | `var zero T` names the zero value of an unknown type. |
| 3 | **Seeding from the first element** | Starting from `s[0]` avoids needing a "smallest possible value" for an unknown type. |

## Hint

Seed with `s[0]`, then scan `s[1:]`.

## Validate

```bash
make verify
```
