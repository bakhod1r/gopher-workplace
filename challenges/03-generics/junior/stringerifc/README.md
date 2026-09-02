# Generics Versus Interfaces

**Level:** junior  
**Topic:** 03-generics

## Context

A logger renders any value that knows how to describe itself. The team is unsure whether to take an interface or a type parameter.

## Task

Implement the stub(s) in [stringerifc.go](stringerifc.go):

1. Implement `Describe`, calling `String()` on each element.
2. `DescribeAny` is provided — compare the two call sites in the tests.
3. Note that a `[]Tag` can be passed to `Describe` but not to `DescribeAny`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Describe([]Tag{"a", "b"})
Output: []string{"tag:a", "tag:b"}
```

**Example 2:**

```
Input:  DescribeAny([]fmt.Stringer{Tag("a")})
Output: []string{"tag:a"}
```

**Example 3:**

```
Input:  Describe([]Tag{})
Output: []string{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Method constraints** | A constraint may be an ordinary interface: `[T fmt.Stringer]` requires the method. |
| 2 | **No slice conversion** | `[]Tag` does not convert to `[]fmt.Stringer` — each element would need boxing. |
| 3 | **When to pick which** | Take an interface for one value; take a type parameter when the caller has a typed collection. |

## Hint

The constraint is just `fmt.Stringer` — interfaces work as constraints unchanged.

## Validate

```bash
make verify
```
