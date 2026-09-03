# Strip Internal Detail

**Level:** senior
**Topic:** 04-error-handling

## Context

An external API must not leak table names or file paths. Known failure classes get a fixed public message; everything else becomes a generic error.

## Task

Implement `Public` in [sanitize.go](sanitize.go):

1. Return `ErrPublicNotFound` for chains containing `errInternalMissing`.
2. Return `ErrPublicInvalid` for chains containing `errInternalParse`.
3. Return `ErrPublicInternal` for anything else, and nil for nil.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Public(fmt.Errorf("table users: %w", errInternalMissing))
Output: ErrPublicNotFound
```

**Example 2:**

```
Input:  Public(errors.New("/etc/secret: denied"))
Output: ErrPublicInternal
```

**Example 3:**

```
Input:  Public(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Information disclosure** | Internal messages are not user-facing. |
| 2 | **Whitelisting** | Only recognised classes get specific answers. |
| 3 | **Boundary responsibility** | Sanitising belongs at the edge, once. |

## Hint

Everything not explicitly recognised must collapse to the generic error — never pass the original message through.

## Validate

```bash
make verify
```
