# Ordering defer relative to acquisition

## Intuition

A `defer cleanup()` placed before the acquisition check runs on every exit path, including failures where nothing was acquired.

## Approach

1. Register the cleanup **after** the resource is acquired.
2. The bug defers `close` before the early return, logging a close with no open.
3. Move the defer below the guard and the `open` step.

## Solution

```go
func Use(ok bool) (log []string) {
	if !ok {
		return
	}
	log = append(log, "open")
	defer func() { log = append(log, "close") }()
	return
}
```

## Walkthrough

Deferring first makes `Use(false)` log `[close]` though nothing opened. Placing the defer after `open` pairs them correctly.

## Pitfalls

- Put `defer resource.Close()` immediately AFTER a successful open, not before.
- A defer above an early error return still executes.
