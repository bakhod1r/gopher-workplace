# Logger

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A service writes log lines to whatever sink the deployment provides — memory in tests, a discard sink in benchmarks.

## Task

Implement the stub(s) in [logger.go](logger.go):

1. Implement `Log` on `*MemLogger` — append the line to `Lines`.
2. Implement `Log` on `Discard` — do nothing.
3. Implement `LogAll`, which logs every message in order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  m := &MemLogger{}; m.Log("a"); m.Lines
Output: ["a"]
```

**Example 2:**

```
Input:  LogAll(&MemLogger{}, []string{"a", "b"})
Output: logger holds ["a", "b"]
```

**Example 3:**

```
Input:  LogAll(Discard{}, []string{"a"})
Output: nothing recorded
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface as a seam for testing** | A memory sink makes output assertable. |
| 2 | **Pointer receiver for state** | `*MemLogger` must mutate; `Discard` needs no state. |
| 3 | **append** | Reused from composite types: growing a slice field. |

## Hint

`Discard.Log` has an empty body — that is the whole implementation.

## Validate

```bash
make verify
```
