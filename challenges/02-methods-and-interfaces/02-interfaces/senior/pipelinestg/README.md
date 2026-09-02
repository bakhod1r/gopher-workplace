# Pipeline Stage

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An ingest pipeline is built from stages that each read a channel and write the next one. A stage that leaks a goroutine takes the process down over days.

## Task

Implement the stub(s) in [pipelinestg.go](pipelinestg.go):

1. Implement `Process` on `DoubleStage` and `DropOddStage`.
2. Implement `RunStage`, which reads `in`, applies the stage, and writes to a returned channel, closing it when `in` closes.
3. Constraint: no goroutine may outlive its input channel — the test asserts the output channel closes and `-race` is clean.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  DoubleStage over [1 2]
Output: [2 4]
```

**Example 2:**

```
Input:  DropOddStage over [1 2 3 4]
Output: [2 4]
```

**Example 3:**

```
Input:  a closed input
Output: the output channel closes too
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pipeline stages** | Each stage owns exactly one output channel and closes it. |
| 2 | **Goroutine lifetime** | Ranging over the input ends when it closes — no leak. |
| 3 | **Interface-driven stages** | Reused: the stage logic is a value, the plumbing is generic. |

## Hint

`defer close(out)` inside the goroutine, then `for v := range in`.

## Validate

```bash
make verify
```
