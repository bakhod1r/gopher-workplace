# Struct Clone Shares Pointer Field

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

Copying `*d` duplicates the struct's fields, but a slice field is just a header
sharing the same backing array. A deep clone must copy the slice too.

## Task

Fix [shallowclone.go](shallowclone.go) so the clone's Tags are independent.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  d := Clone(src); c.Tags[0] = "X"
Output: src.Tags unchanged
```

**Example 2:**

```
Input:  Clone(&Doc{})
Output: independent empty Doc
```

**Example 3:**

```
Input:  len(Clone(d).Tags)
Output: len(d.Tags)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shallow struct copy** | `*d` shares reference-typed fields. |
| 2 | **Copy the slice** | `append([]string(nil), d.Tags...)`. |
| 3 | **Deep clone** | Independent backing arrays. |

## Hint

Copy the slice field too: `cp := *d; cp.Tags = append([]string(nil), d.Tags...); return &cp`.

## Validate

```bash
make verify
```
