# Generic Test Helper

**Level:** junior  
**Topic:** 03-generics

## Context

Every test file repeats the same four-line comparison. One typed helper removes the noise without hiding the values.

## Task

Implement the stub(s) in [assertequalhelper.go](assertequalhelper.go):

1. Implement `Equal`, reporting the comparison result.
2. Call `t.Helper()` so failures point at the caller's line, not this file.
3. Record the failure with `t.Errorf` in the form `got X, want Y`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Equal(t, 1, 1)
Output: true, no failure
```

**Example 2:**

```
Input:  Equal(t, 1, 2)
Output: false, one failure recorded
```

**Example 3:**

```
Input:  Equal(t, "a", "a")
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`testing.TB`** | The interface both `*testing.T` and `*testing.B` satisfy — take it, not `*testing.T`. |
| 2 | **`t.Helper()`** | Marks this function as a helper so reported line numbers point at the caller. |
| 3 | **Typed assertions** | `comparable` gives a compile error when the types differ, which a reflect-based helper would not. |

## Hint

`t.Helper()` first, then compare.

## Validate

```bash
make verify
```
