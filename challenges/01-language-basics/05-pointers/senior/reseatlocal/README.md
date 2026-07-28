# Reassigning a Local Pointer

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Assigning to `pp` changes the local copy of the double pointer. To reseat the
caller's pointer you must dereference once: `*pp = q`.

## Task

Fix [reseatlocal.go](reseatlocal.go).

Do **not** change the function signature or the tests.

## Examples

```go
p := &a; Reseat(&p, &b) // p == &b
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Levels of indirection** | `*pp` is the caller's pointer. |
| 2 | **Local copy** | `pp` itself is a copy of `&p`. |
| 3 | **Dereference once** | `*pp = q`. |

## Hint

Dereference the double pointer: `*pp = q`.

## Validate

```bash
make verify
```
