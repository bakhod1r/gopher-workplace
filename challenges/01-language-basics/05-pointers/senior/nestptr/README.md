# Nested Pointer Map Init

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

The inner map `m[group]` is nil until created; writing `m[group][key]` panics.
Lazily initialise the inner map before the write.

## Task

Fix [nestptr.go](nestptr.go).

Do **not** change the function signature or the tests.

## Examples

```go
Add(m, "g", "a") twice // m["g"]["a"] == 2
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nested maps** | The inner map starts nil. |
| 2 | **Lazy init** | Create `m[group]` on first use. |
| 3 | **Nil map write** | Writing to nil panics. |

## Hint

Init the inner map: `if m[group] == nil { m[group] = map[string]int{} }`, then `m[group][key]++`.

## Validate

```bash
make verify
```
