# Two-Value Result

**Level:** junior  
**Topic:** 03-generics

## Context

A batch job records, per item, whether it worked and why not — and the item type differs per job.

## Task

Implement the stub(s) in [resultgen.go](resultgen.go):

1. Implement `Ok`, `Fail`, `Unwrap`, and `Reason`.
2. A successful result reports an empty reason; a failed one reports the zero value from `Unwrap`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Ok(5).Unwrap()
Output: 5, true
```

**Example 2:**

```
Input:  Fail[int]("bad").Unwrap()
Output: 0, false
```

**Example 3:**

```
Input:  Fail[int]("bad").Reason()
Output: "bad"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Mixing generic and concrete fields** | The value is `T`; the reason is always a `string`. |
| 2 | **Explicit instantiation on failure** | `Fail[int]("bad")` must name `T` — the reason says nothing about the value type. |
| 3 | **Value receivers** | Neither method mutates, so both take the value. |

## Hint

`Fail[int]("bad")` needs its type argument spelled out.

## Validate

```bash
make verify
```
