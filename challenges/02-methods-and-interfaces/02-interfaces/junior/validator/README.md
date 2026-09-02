# Validator

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A signup form runs several independent rules over the submitted value.

## Task

Implement the stub(s) in [validator.go](validator.go):

1. Implement `Validate` on `NotEmpty` and `MaxLen`.
2. Implement `ValidateAll`, which runs every rule and returns the first error, or nil.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NotEmpty{}.Validate("a")
Output: nil
```

**Example 2:**

```
Input:  NotEmpty{}.Validate("")
Output: error "empty"
```

**Example 3:**

```
Input:  ValidateAll([]Validator{NotEmpty{}, MaxLen{N: 2}}, "abc")
Output: error "too long"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Rule as an interface** | Rules compose without a giant `if` chain. |
| 2 | **errors.New sentinels** | Reused from error handling: comparable error values. |
| 3 | **len on strings** | Reused: byte length as the size rule. |

## Hint

`ValidateAll` returns as soon as a rule returns non-nil.

## Validate

```bash
make verify
```
