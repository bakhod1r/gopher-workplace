# Update Through Fetched Pointer

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

Fetching a pointer from a map with comma-ok lets you mutate the pointee and
report presence.

## Task

Implement `SetName` in [setfield.go](setfield.go).

Do **not** change the function signature or the tests.

## Examples

```go
SetName(m, 1, "new") // m[1].Name = "new", true
SetName(m, 99, ..)   // false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **comma-ok fetch** | `u, ok := m[id]`. |
| 2 | **Mutate pointee** | `u.Name = name`. |
| 3 | **Report presence** | Return ok. |

## Hint

`u, ok := m[id]; if !ok { return false }; u.Name = name; return true`.

## Validate

```bash
make verify
```
