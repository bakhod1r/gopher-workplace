# The Assertion That Always Passes

**Level:** staff  
**Topic:** 03-generics

## Context

A shared `AssertEqual` helper is used as a guard: callers stop dereferencing a result when it returns false. It never returns false, so a mismatch is followed by a nil dereference inside the caller's next line.

## Task

Fix the single planted bug in [asserthelperpassbug.go](asserthelperpassbug.go):

1. Find and fix the single bug so the helper reports failure for unequal values.
2. Equal values must still report success and log nothing.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  AssertEqual(t, 1, 1)
Output: true, nothing logged
```

**Example 2:**

```
Input:  AssertEqual(t, 1, 2)
Output: false, one message logged
```

**Example 3:**

```
Input:  AssertEqual(t, "a", "b")
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Test helpers must fail loudly** | A helper that reports success on a mismatch turns a red suite green. |
| 2 | **Report and return are two jobs** | Logging the mismatch does not tell the caller to stop. |
| 3 | **Guard-shaped helpers** | A boolean result is a control-flow contract: `false` must mean "do not continue". |

## Hint

Follow the code path taken when `got != want`.

## Validate

```bash
make verify
```
