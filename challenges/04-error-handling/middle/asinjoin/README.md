# Find A Type Inside A Join

**Level:** middle
**Topic:** 04-error-handling

## Context

A bulk validator returns joined failures. The handler needs the first field-level error among them so it can point at the offending input.

## Task

Implement `FirstField` in [asinjoin.go](asinjoin.go):

1. Return the first `*FieldError` found in the error, and true.
2. Search inside joined errors as well as wrapping chains.
3. Return `nil, false` when no `*FieldError` is present.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FirstField(errors.Join(ErrOther, &FieldError{Name: "age"}))
Output: the *FieldError, true
```

**Example 2:**

```
Input:  FirstField(ErrOther)
Output: nil, false
```

**Example 3:**

```
Input:  FirstField(nil)
Output: nil, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **errors.As over trees** | `As` searches joined branches, not just chains. |
| 2 | **Typed targets** | The target is a pointer to the pointer type. |
| 3 | **Depth-first order** | The first match in tree order wins. |

## Hint

`errors.As` already understands both unwrapping shapes — no manual traversal is needed.

## Validate

```bash
make verify
```
