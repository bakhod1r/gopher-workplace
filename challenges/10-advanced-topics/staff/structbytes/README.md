# A Byte View Of A Struct, Only When It Is Safe

**Level:** staff
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A wire encoder writes frames by copying each field into a buffer. Profiling says the copy is most of the send path, and every frame is scalars only.

## Task

Implement [structbytes.go](structbytes.go):

1. Return a byte view of `*p`, of exactly `unsafe.Sizeof(*p)` bytes.
2. Return false for a nil pointer, and for any struct type that contains a pointer-shaped field.
3. Zero allocations; the view's capacity must equal its length.

Replace the stub body in [structbytes.go](structbytes.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Bytes(&Frame{})
Output: a 16-byte view, true
```

**Example 2:**

```
Input:  Bytes(nil)
Output: nil, false
```

**Example 3:**

```
Input:  a struct containing a *int
Output: nil, false
```

_Explanation:_ A byte view would expose an address.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pointer-free is the precondition** | Bytes of a pointer are an address, meaningless to a peer and invisible to the collector. |
| 2 | **unsafe.Slice over a struct pointer** | Length and capacity both come from `Sizeof`. |
| 3 | **Padding is in the view** | The bytes include the struct's padding, whose contents are unspecified. |
| 4 | **Endianness and layout** | The view is the machine's layout, not a portable wire format. |

## Hint

`hasPointers` is written for you. Guard, check, then build the view.

## Validate

```bash
make verify
```
