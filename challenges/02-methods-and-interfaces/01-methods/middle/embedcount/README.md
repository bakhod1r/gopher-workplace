# Embedded Count

**Level:** middle
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Embedding is a way to reuse state and methods. A `Job` embeds a `Tracker` to
count how many times it runs.

## Task

Implement `Run` on `*Job` in [embedcount.go](embedcount.go):

1. Call the promoted `Inc` method.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  j := Job{Name: "task"}; j.Run()
Output: j.Count == 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Promoted fields and methods** | `j.Count` and `j.Inc()` are directly accessible. |
| 2 | **Pointer receiver promotion** | `Inc` needs `*Tracker`. `Run` has `*Job`, which provides a pointer to the embedded `Tracker`. |

## Hint

`j.Inc()` — that's it.

## Validate

```bash
make verify
```
