# No Operator Overloading

**Level:** middle  
**Topic:** 03-generics

## Context

Money values must not be added with a bare `+` — the currency has to match. The type provides `Plus`, and generic code must use it.

## Task

Implement the stub(s) in [addviamethodgen.go](addviamethodgen.go):

1. Implement `SumAll`, folding with each value's `Plus` method.
2. Start from the zero value of `T`.
3. Understand why `+` is unavailable for a struct type parameter.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SumAll([]Money{{2},{3}})
Output: Money{5}
```

**Example 2:**

```
Input:  SumAll([]Money{})
Output: Money{0}
```

**Example 3:**

```
Input:  SumAll([]Money{{1}})
Output: Money{1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **No operator overloading** | Go will never let a struct define `+`; a method is the only option. |
| 2 | **Self-referential constraint** | `Adder[T]` requires `Plus(T) T` — the same shape as `Lesser[T]`. |
| 3 | **Zero value as identity** | The fold starts from `var out T`, which must be the additive identity. |

## Hint

`out = out.Plus(v)` — the method is the operator.

## Validate

```bash
make verify
```
