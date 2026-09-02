# Lazy Load

## Intuition

`sync.Once` does two jobs: it serialises the first call, and it publishes everything that call wrote. That second half is why later readers can touch `l.value` without a lock.

## Approach

1. `Get` wraps the build in `l.once.Do`.
2. Inside, assign `l.value` and set the `built` flag.
3. Return `l.value` unconditionally — after `Do` returns, the write is visible.
4. `Built` reads the atomic flag so it is safe to call concurrently.

## Solution

```go
func (l *Lazy) Get() string {
	l.once.Do(func() {
		l.value = l.builder.Build()
		l.built.Store(true)
	})
	return l.value
}

func (l *Lazy) Built() bool { return l.built.Load() }
```

## Walkthrough

With 100 concurrent `Get` calls, 99 block inside `Do` until the leader finishes. When they resume, the memory model guarantees they see the leader's `l.value` write.

## Pitfalls

- A `if !built { build() }` check without synchronisation — two goroutines can both see false.
- Copying a `Lazy` value after use: `sync.Once` must not be copied.
- Using a plain `bool` for `built` and reading it concurrently, which `-race` flags.
