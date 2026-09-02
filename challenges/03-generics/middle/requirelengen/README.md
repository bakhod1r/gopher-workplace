# Fatal Assertion Helper

**Level:** middle  
**Topic:** 03-generics

## Context

Tests that index into a result crash with an unhelpful panic when the result is shorter than expected.

## Task

Implement the stub(s) in [requirelengen.go](requirelengen.go):

1. Implement `RequireLen`, failing the test immediately on a wrong length.
2. Use `Fatalf`, not `Errorf` — the point is to stop before the indexing that follows.
3. Include both lengths and the slice in the message.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RequireLen(t, []int{1}, 1)
Output: no failure
```

**Example 2:**

```
Input:  RequireLen(t, []int{}, 1)
Output: test stops
```

**Example 3:**

```
Input:  after RequireLen
Output: indexing is safe
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fatal versus error** | `Fatalf` stops the goroutine; `Errorf` records and continues. |
| 2 | **Guarding before indexing** | A fatal check turns a panic into a readable failure. |
| 3 | **`T any` is enough** | Nothing is compared, so the elements need no constraint. |

## Hint

`Fatalf` here, unlike the `Errorf` in the equality helper — the difference is deliberate.

## Validate

```bash
make verify
```
