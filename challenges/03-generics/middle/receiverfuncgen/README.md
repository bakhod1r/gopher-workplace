# Functions That Act Like Methods

**Level:** middle  
**Topic:** 03-generics

## Context

A wrapper type needs both an in-place update and a type-changing conversion. Only one of them can be a method.

## Task

Implement the stub(s) in [receiverfuncgen.go](receiverfuncgen.go):

1. Implement `Update`, `Convert`, and `Get`.
2. `Update` mutates the box in place; `Convert` returns a new box of a different element type.
3. Note which one had to be a function and why.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Update(b, double); b.Get()
Output: the doubled value
```

**Example 2:**

```
Input:  Convert(b, itoa).Get()
Output: the string form
```

**Example 3:**

```
Input:  Convert leaves b alone
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method-shaped functions** | Taking `*Box[T]` first reads almost like a method call. |
| 2 | **Type-changing operations** | Anything producing `Box[U]` must be a function. |
| 3 | **Mutation versus construction** | `Update` writes through the pointer; `Convert` allocates. |

## Hint

`Update` could have been a method; `Convert` could not.

## Validate

```bash
make verify
```
