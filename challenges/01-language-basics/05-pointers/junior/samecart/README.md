# Same Struct Instance

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Comparing struct POINTERS tests instance identity; two structs with equal
fields are still different instances.

## Task

Implement `Same` in [samecart.go](samecart.go).

Do **not** change the function signature or the tests.

## Examples

```go
Same(c, c) // => true
Same(c, d) // => false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer identity** | `a == b` compares addresses. |
| 2 | **Instance vs value** | Equal fields, different instances. |
| 3 | **Struct pointers** | `*Cart` addresses. |

## Hint

`return a == b`.

## Validate

```bash
make verify
```
