# All Checks Passed

**Level:** junior
**Topic:** 04-error-handling

## Context

A release gate runs a list of checks and only proceeds when every one of them succeeded.

## Task

Implement `AllNil` in [allnil.go](allnil.go):

1. Return `true` when every entry is nil.
2. Return `false` as soon as any entry is non-nil.
3. Return `true` for an empty or nil slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  AllNil([]error{nil, nil})
Output: true
```

**Example 2:**

```
Input:  AllNil([]error{nil, ErrCheck})
Output: false
```

**Example 3:**

```
Input:  AllNil(nil)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Vacuous truth** | Nothing failed in an empty slice, so the answer is true. |
| 2 | **Early exit** | One failure settles the question. |
| 3 | **Loop then return** | The success answer comes after the loop. |

## Hint

Return `false` inside the loop, `true` after it — not the other way around.

## Validate

```bash
make verify
```
