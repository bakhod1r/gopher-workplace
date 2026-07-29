# Nil Check Before Unsafe Read

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _nil-pointer-dereference_

## Context

Even with unsafe pointers, a nil check must PRECEDE the dereference. Reading
`*(*int)(p)` first panics when p is nil.

## Task

Fix [nilunsafe.go](nilunsafe.go) so a nil pointer returns the default without panicking.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ReadOr(nil, 7)
Output: 7
```

**Example 2:**

```
Input:  x := 5; ReadOr(unsafe.Pointer(&x), 7)
Output: 5
```

**Example 3:**

```
Input:  ReadOr(nil, 0)
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Guard before deref** | Check nil first, even for unsafe. |
| 2 | **Order of operations** | The panic is at the read. |
| 3 | **Default on nil** | Return def. |

## Hint

Check nil first: `if p == nil { return def }; return *(*int)(p)`.

## Validate

```bash
make verify
```
