# Once Initialization

## Intuition

Lazy initialization has a race baked in: "check if it is built, and if not,
build it" is two steps, and two goroutines can both pass the check.
`sync.Once` collapses those steps into one atomic decision and — crucially —
makes every other caller *wait* until the winner finishes.

## Approach

1. Put the whole initialization inside the `Do` closure.
2. Read the field after `Do` returns; by then it is guaranteed written.

## Solution

```go
func (l *LazyData) Get() string {
	l.once.Do(func() {
		l.data = l.init()
	})
	return l.data
}
```

## Walkthrough

Ten goroutines call `Get`. Exactly one enters the closure; the other nine block
inside `Do` until it returns. The test's `atomic.AddInt32` therefore reaches 1,
and every goroutine reads `"safe"`.

The blocking is not incidental — it is what gives the happens-before edge. The
write to `l.data` inside `Do` is guaranteed visible to every later `Do` caller,
so reading `l.data` after `Do` needs no extra lock.

## Pitfalls

- **Assigning outside the closure.** `l.once.Do(l.init)` does not compile
  (`init` returns a value), and computing outside `Do` reintroduces the race.
- **Copying `LazyData`.** A copied `sync.Once` has a fresh state and will run
  the initializer again; `go vet` flags this as a lock copy.
- **A plain `if l.data == ""` guard.** Racy, and wrong when the real value is
  the empty string.

## Run it with `-race`

`go test -race` is what turns this from "looks fine" to "proven". The unguarded
version fails the race detector on the concurrent write to `l.data`.
