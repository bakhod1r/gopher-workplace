# Deferred Panic Masks Original

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

If a deferred function panics during unwinding, its panic REPLACES the one in
flight; the outer recover then sees the cleanup's value, not the original. A
cleanup must not panic (or must recover its own).

## Task

Fix [maskpanic.go](maskpanic.go) so the original panic value survives.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FirstPanic(func(){ panic("original") })
Output: "original"
```

**Example 2:**

```
Input:  deferred cleanup does not mask
Output: true
```

**Example 3:**

```
Input:  FirstPanic reports first panic
Output: original
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Panic during unwind** | A deferred panic supersedes the active one. |
| 2 | **Cleanup discipline** | Cleanups must not panic. |
| 3 | **recover sees the latest** | The outer recover captures whatever panic is current. |

## Hint

Remove the panicking cleanup (or make it recover its own panic) so the original value reaches the outer recover.

## Validate

```bash
make verify
```
