# Map Needs A Function

**Level:** junior  
**Topic:** 03-generics

## Context

The team wants `data.Map(toString)` to work like `data.Filter(ok)`. Go refuses, and the reason is worth understanding.

## Task

Implement the stub(s) in [mapmethodlesson.go](mapmethodlesson.go):

1. Implement `MapSlice` as a **function** taking `Slice[T]` and returning `Slice[U]`.
2. Implement `Each` as a **method** that calls `f` for every element in order.
3. Note why `Map` cannot be a method: it would need a type parameter of its own.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MapSlice(Slice[int]{1,2}, double)
Output: Slice[int]{2,4}
```

**Example 2:**

```
Input:  MapSlice(Slice[int]{1,2}, itoa)
Output: Slice[string]{"1","2"}
```

**Example 3:**

```
Input:  Slice[int]{1,2}.Each(collect)
Output: collect called with 1, then 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Methods take no new type parameters** | Go allows type parameters on the type, never extra ones on its methods. |
| 2 | **Where the extra parameter can live** | A function may declare `[T, U any]`; a method may only use the type's own parameters. |
| 3 | **Design consequence** | This is why the stdlib exposes `slices.Map`-style helpers as functions. |

## Hint

`Each` introduces no new type, so it fits as a method; `Map` introduces `U`, so it cannot.

## Validate

```bash
make verify
```
