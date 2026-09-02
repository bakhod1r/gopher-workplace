# Unsafe Slice And String

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A parser converted `[]byte` to `string` on every field. The copies dominated the profile, and the fix uses the no-copy conversions — with strict rules about mutation.

## Task

Implement the stub(s) in [unsafeslice.go](unsafeslice.go):

1. Implement `BytesToString` with `unsafe.String` (no copy) and `StringToBytes` with `unsafe.Slice` (no copy).
2. Implement `SafeString`, the copying version.
3. Constraint: `BytesToString` must allocate zero times, and the docs must state that mutating the source afterwards changes the string.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  BytesToString([]byte("abc"))
Output: "abc", no allocation
```

**Example 2:**

```
Input:  mutating the source afterwards
Output: the unsafe string changes too
```

**Example 3:**

```
Input:  SafeString on the same input
Output: an independent copy
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **unsafe.String / unsafe.Slice** | The supported no-copy conversions between bytes and strings. |
| 2 | **Immutability invariant** | Strings are assumed immutable by the runtime; breaking that is undefined behaviour. |
| 3 | **Empty-input edge case** | A zero-length slice may have a nil data pointer — handle it explicitly. |

## Hint

`unsafe.String(&b[0], len(b))` panics on an empty slice — return `""` first.

## Validate

```bash
make verify
```
