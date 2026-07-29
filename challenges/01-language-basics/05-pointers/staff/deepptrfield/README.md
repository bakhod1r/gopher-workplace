# Copy Shares Pointer Field

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Copying `*b` duplicates the pointer field's VALUE (the address), so both boxes
share one pointee. A deep clone must allocate a new int and copy the value.

## Task

Fix [deepptrfield.go](deepptrfield.go) so the clone's pointer is independent.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  c := Clone(b); *c.P = 99
Output: *b.P unchanged
```

**Example 2:**

```
Input:  Clone(&Box{P: &x})
Output: independent pointee
```

**Example 3:**

```
Input:  c.P != b.P
Output: true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shallow copy shares pointers** | `*b` copies the address. |
| 2 | **Deep copy the pointee** | Allocate a new int, copy the value. |
| 3 | **Independence** | Distinct pointees. |

## Hint

Deep-copy the pointee: `cp := *b; v := *b.P; cp.P = &v; return &cp`.

## Validate

```bash
make verify
```
