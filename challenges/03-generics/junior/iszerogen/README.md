# Is Zero

**Level:** junior  
**Topic:** 03-generics

## Context

A form validator marks unset fields. "Unset" means the zero value, whatever the field's type happens to be.

## Task

Implement the stub(s) in [iszerogen.go](iszerogen.go):

1. Implement `IsZero`, reporting whether `v` equals the zero value of `T`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IsZero(0)
Output: true
```

**Example 2:**

```
Input:  IsZero("")
Output: true
```

**Example 3:**

```
Input:  IsZero(3)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Zero value of `T`** | `var zero T` is the only way to name the zero value of an unknown type. |
| 2 | **The `comparable` constraint** | `comparable` is what lets you use `==` and `!=` on a type parameter. |
| 3 | **Zero values per type** | Reused from language basics: `0`, `""`, `false`, and the nil pointer are all zero values. |

## Hint

Declare `var zero T` and compare with `==`.

## Validate

```bash
make verify
```
