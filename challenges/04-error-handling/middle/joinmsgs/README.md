# List Joined Messages

**Level:** middle
**Topic:** 04-error-handling

## Context

An API renders each validation failure as its own JSON entry, so the combined error must be taken apart into plain strings.

## Task

Implement `Lines` in [joinmsgs.go](joinmsgs.go):

1. Return one message per error inside a joined error.
2. Return a single-element slice for a non-joined error.
3. Return nil for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Lines(errors.Join(ErrA, ErrB))
Output: ["a" "b"]
```

**Example 2:**

```
Input:  Lines(ErrA)
Output: ["a"]
```

**Example 3:**

```
Input:  Lines(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Unwrap() []error** | The multi-error shape exposes siblings. |
| 2 | **Rendering vs matching** | Presentation code needs strings, not sentinels. |
| 3 | **Single-error fallback** | A plain error is a list of one. |

## Hint

Splitting `err.Error()` on newlines happens to work for `errors.Join` today — asking the type for its members does not depend on that.

## Validate

```bash
make verify
```
