# String-Like Identifiers

**Level:** middle  
**Topic:** 03-generics

## Context

Every entity has its own ID type — `UserID`, `OrderID` — so one cannot be passed where another is expected. Validation should still be shared.

## Task

Implement the stub(s) in [idvalidate.go](idvalidate.go):

1. Implement `ValidID`, reporting whether `v` is non-empty and starts with `prefix`.
2. The `~string` constraint is what lets the named ID types through.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ValidID(UserID("u_1"), "u_")
Output: true
```

**Example 2:**

```
Input:  ValidID(UserID("x_1"), "u_")
Output: false
```

**Example 3:**

```
Input:  ValidID(UserID(""), "u_")
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The `~` token** | `~int` admits every type whose underlying type is `int`. |
| 2 | **Converting to use stdlib** | `string(v)` is needed because `strings` functions take a plain `string`. |
| 3 | **Zero-cost conversion** | Converting a `~string` type parameter to `string` copies no bytes at run time. |

## Hint

Convert once with `string(v)`, then use the `strings` package normally.

## Validate

```bash
make verify
```
