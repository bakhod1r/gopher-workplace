# Element Pointer After Append

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

The append reallocates (len == cap), so `p` (taken before) points at the OLD
array. Writing through it doesn't touch the returned slice; re-take the address
or index directly after the append.

## Task

Fix [elemstale.go](elemstale.go) so the first element of the result is 42.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FirstOf([1 2])
Output: [42 2 0]
```

**Example 2:**

```
Input:  result length
Output: 3
```

**Example 3:**

```
Input:  s[0] after call
Output: 42
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Reallocation on append** | len==cap forces a new array. |
| 2 | **Stale element pointer** | Old `&s[0]` is detached. |
| 3 | **Re-address** | Index or re-take after the append. |

## Hint

Write through the current slice: `s[0] = 42` (or re-take `p = &s[0]`).

## Validate

```bash
make verify
```
