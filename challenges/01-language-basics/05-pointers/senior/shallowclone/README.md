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

```go
Clone(d); c.Tags[0]="X" // d.Tags unchanged
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
