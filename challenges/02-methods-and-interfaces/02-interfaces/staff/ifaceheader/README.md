# Interface Header

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A nil check on an error kept passing when the caller was sure the error was nil. The fix requires seeing what an interface value actually holds.

## Task

Implement the stub(s) in [ifaceheader.go](ifaceheader.go):

1. Implement `Words`, returning the two words of an `any` — the type pointer and the data pointer.
2. Implement `IsTypedNil`, reporting whether an interface holds a non-nil type with a nil data word.
3. Implement `Classify`, returning `"nil"`, `"typed-nil"`, or `"value"`.
4. Constraint: no `reflect` — read the header directly with `unsafe`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Classify(nil)
Output: "nil"
```

**Example 2:**

```
Input:  Classify((*int)(nil))
Output: "typed-nil"
```

**Example 3:**

```
Input:  Classify(42)
Output: "value"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **eface layout** | An `any` is two words: a `*_type` and a data pointer. |
| 2 | **Typed nil** | A non-nil type word with a nil data word — non-nil as an interface. |
| 3 | **unsafe.Pointer rules** | Reused: converting between a pointer and a struct of the same layout. |

## Hint

`*(*[2]uintptr)(unsafe.Pointer(&v))` gives the type word and the data word of an `any`.

## Validate

```bash
make verify
```
