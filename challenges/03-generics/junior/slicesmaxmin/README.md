# Max And Min From Stdlib

**Level:** junior  
**Topic:** 03-generics

## Context

A gauge reads the extremes of a sample window. The window is empty until the first scrape lands.

## Task

Implement the stub(s) in [slicesmaxmin.go](slicesmaxmin.go):

1. Implement `Peak` and `Floor` using `slices.Max` and `slices.Min`.
2. Both must return the zero value and `false` for an empty slice — the stdlib helpers panic there.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Peak([]int{1, 9})
Output: 9, true
```

**Example 2:**

```
Input:  Floor([]int{1, 9})
Output: 1, true
```

**Example 3:**

```
Input:  Peak([]int{})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Panicking stdlib helpers** | `slices.Max` and `slices.Min` panic on an empty slice — read the docs before delegating. |
| 2 | **Wrapping to make total** | Adding a guard turns a panicking helper into a checked API. |
| 3 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<` and `>`. |

## Hint

Guard the empty case yourself — the stdlib will not.

## Validate

```bash
make verify
```
