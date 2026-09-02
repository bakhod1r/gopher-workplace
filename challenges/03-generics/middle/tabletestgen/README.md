# Generic Table Runner

**Level:** middle  
**Topic:** 03-generics

## Context

Every test file repeats the same table-driven loop. One typed runner removes the boilerplate without hiding the failure output.

## Task

Implement the stub(s) in [tabletestgen.go](tabletestgen.go):

1. Implement `Run`, executing each case as a subtest named after the case.
2. Report failures with the input, the actual value, and the expected value.
3. Call `t.Helper()` so failures point at the caller.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Run(t, "double", cases, double)
Output: one subtest per case
```

**Example 2:**

```
Input:  a failing case
Output: reports input, got and want
```

**Example 3:**

```
Input:  no cases
Output: no subtests
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two comparable parameters** | Input and expectation may differ in type; both need `==`. |
| 2 | **Subtests** | `t.Run` gives each case its own name and failure line. |
| 3 | **Helpers in tests** | `t.Helper()` keeps reported lines pointing at the table, not the runner. |

## Hint

The type parameters make a mismatched table a compile error, not a run-time surprise.

## Validate

```bash
make verify
```
