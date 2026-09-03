# Collect Every Typed Error

**Level:** staff
**Topic:** 04-error-handling

## Context

A validation report needs every field-level error in the tree, not just the first one `errors.As` happens to find.

## Task

Implement `All` in [asall.go](asall.go):

1. Return every `*FieldError` in the tree, depth first, left to right.
2. Search wrapped and joined children alike.
3. Return nil when there are none.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  All(errors.Join(&FieldError{"a"}, &FieldError{"b"}))
Output: both
```

**Example 2:**

```
Input:  All(ErrOther)
Output: nil
```

**Example 3:**

```
Input:  All(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Beyond errors.As** | `As` stops at the first match. |
| 2 | **Manual traversal** | Collecting requires visiting everything. |
| 3 | **Type assertion per node** | The node itself may be a match. |

## Hint

`errors.As` short-circuits by design — collecting all matches means walking the tree yourself.

## Validate

```bash
make verify
```
