# 06 Concurrency: Topics to Master

Concurrency is split into **5 sections**. Each section holds 4 levels
(`junior/ middle/ senior/ staff`) with 30 puzzles per level — 120 per section,
600 in total.

| # | Section | What it covers |
|---|---------|----------------|
| 01 | Goroutines | `go`, `sync.WaitGroup`, loop-variable capture, `defer wg.Done()`, goroutine lifetime, leaks, panic isolation, stack growth, scheduler basics |
| 02 | Channels | buffered vs unbuffered, direction (`chan<-`, `<-chan`), `close`, `range`, comma-ok, nil channels, `select`, `default`, timeouts, `time.After`/`Ticker`/`Timer`, channel internals |
| 03 | Sync and Atomics | `Mutex`, `RWMutex`, `Cond`, `Once`, `Pool`, `Map`, `sync/atomic`, the Go memory model, lock-free structures |
| 04 | Context and Cancellation | `WithCancel`/`WithTimeout`/`WithDeadline`/`WithValue`, propagation, `ctx.Err()`, cleanup ordering, `errgroup` |
| 05 | Patterns and Pitfalls | pipeline, fan-in/fan-out, worker pool, semaphore, rate limiting, deadlock, livelock, starvation, data races, anti-patterns, testing under `-race` |

## Level profiles inside a section

| Level | Focus |
|-------|-------|
| 🟢 junior | Use the primitive correctly on a small, deterministic task. |
| 🔵 middle | Coordinate many goroutines; ordering, ownership, and shutdown. |
| 🟠 senior | Find the planted defect: leak, deadlock, race, lost wakeup. |
| 🔴 staff | Memory model, `-race` cleanliness, and a CPU/time ceiling. |

## Rules for concurrency puzzles

- Every test must pass under `go test -race` and within `-timeout 30s`.
- `time.Sleep` is **not** a synchronisation primitive — no test may depend on
  it for correctness. Use `WaitGroup`, channels, or `context`.
- Tests must be deterministic: no flaky assertions on scheduling order.

## Dropped from the roadmap tree

`00-introduction` (not a coding puzzle), `14-performance-tuning` (belongs to
`11-performance-engineering`), `23-concurrency-in-stdlib` and
`24-primitives-decision-guide` (reference material, folded into the sections
above), `25-modern-features` (belongs to `17-modern-language-features`).
