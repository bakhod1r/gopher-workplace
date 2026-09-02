# Value Or Default

**Level:** junior
**Topic:** 04-error-handling

## Context

A settings reader falls back to a built-in default whenever the lookup failed. Callers never see the failure.

## Task

Implement `OrDefault` in [ordefault.go](ordefault.go):

1. Return `v` when `err` is nil.
2. Return `def` when `err` is non-nil.
3. An empty `v` with a nil error is still returned as-is.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  OrDefault("8080", nil, "80")
Output: "8080"
```

**Example 2:**

```
Input:  OrDefault("", ErrMissing, "80")
Output: "80"
```

**Example 3:**

```
Input:  OrDefault("", nil, "80")
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fallback on error** | The error, not the value, selects the branch. |
| 2 | **Deliberate empty values** | A nil error means the empty value was intended. |
| 3 | **Swallowing errors** | Only safe when a default is genuinely correct. |

## Hint

One test case pairs an empty value with a nil error — it must not fall back.

## Validate

```bash
make verify
```
