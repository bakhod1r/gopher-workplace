# Generic Box

**Level:** junior  
**Topic:** 03-generics

## Context

A lazily-loaded setting must distinguish "never loaded" from "loaded and happens to be zero".

## Task

Implement the stub(s) in [boxgen.go](boxgen.go):

1. Implement `Set`, storing the value and marking the box filled.
2. Implement `Get`, returning the stored value and `true`, or the zero value and `false` when nothing was stored.
3. Storing a zero value still counts as filled.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Box[int]{}.Get()
Output: 0, false
```

**Example 2:**

```
Input:  Set(0); Get()
Output: 0, true
```

**Example 3:**

```
Input:  Set(5); Get()
Output: 5, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Generic types** | `type Stack[T any] struct { ... }` parameterises the type itself, not just a function. |
| 2 | **Zero value of `T`** | `var zero T` names the zero value of an unknown type. |
| 3 | **Presence flags** | A separate `bool` field is how a struct records "was this ever set?". |

## Hint

Storing `0` must still make `Get` report `true` — that is what the `filled` field is for.

## Validate

```bash
make verify
```
