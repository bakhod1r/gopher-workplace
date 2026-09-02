# Email Shape

**Level:** junior
**Topic:** 04-error-handling

## Context

A newsletter form does a cheap sanity check before queueing an address. Full validation happens later; this only catches obvious typos.

## Task

Implement `ValidEmail` in [validemail.go](validemail.go):

1. Return `ErrNoAt` when `s` contains no `@`.
2. Return `ErrEmptyPart` when the text before or after the `@` is empty.
3. Return nil when both parts are non-empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ValidEmail("a@b.com")
Output: nil
```

**Example 2:**

```
Input:  ValidEmail("abc")
Output: ErrNoAt
```

**Example 3:**

```
Input:  ValidEmail("@b.com")
Output: ErrEmptyPart
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **strings.Cut** | Splits once and reports whether the separator was found. |
| 2 | **Ordered guards** | Structure is checked before content. |
| 3 | **Distinct sentinels** | Each rejection reason is its own value. |

## Hint

`strings.Cut` gives you both parts and a found flag in one call — the two failures fall out of it directly.

## Validate

```bash
make verify
```
