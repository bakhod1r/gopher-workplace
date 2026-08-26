# Pool Worker

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A concurrent task executor uses a worker pool to limit the number of active
goroutines. The pool state and behaviour are encapsulated in methods.

## Task

Implement `Start` on `*Pool` in [poolworker.go](poolworker.go):

1. Launch `p.Count` worker goroutines.
2. In each worker, use `for task := range p.Tasks { task() }`.
3. Use `p.wg` to track the workers (Add before launch, Done when finished).

Do **not** change the function signatures or the tests.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Worker pools** | Long-running goroutines reading from a shared channel. |
| 2 | **WaitGroup in methods** | Encapsulating synchronization state inside the struct. |

## Validate

```bash
make verify
```
