# Method Sets

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A registry accepts anything that satisfies its interface, and half the submitted types silently fail to compile — or fail at runtime.

## Task

Implement the stub(s) in [methodset.go](methodset.go):

1. Implement `Name` on `Value` (value receiver) and `Name` on `*Pointer` (pointer receiver).
2. Implement `Rename` on `*Pointer`.
3. Implement `Names`, which returns the `Name()` of every element in a `[]Named`.
4. Implement `Satisfies`, which reports whether the value passed satisfies `Renamer`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Names([]Named{Value{N: "a"}, &Pointer{N: "b"}})
Output: ["a", "b"]
```

**Example 2:**

```
Input:  Satisfies(&Pointer{})
Output: true
```

**Example 3:**

```
Input:  Satisfies(Pointer{})
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method sets** | A `T` value has only value-receiver methods; `*T` has both. |
| 2 | **Addressability** | Calling a pointer method on an addressable variable is sugar; interface storage is not. |
| 3 | **Interface assertion** | Reused: comma-ok assertion to an interface type. |

## Hint

A `T` value in an interface loses access to `*T`'s methods, because the copy has no address.

## Validate

```bash
make verify
```
