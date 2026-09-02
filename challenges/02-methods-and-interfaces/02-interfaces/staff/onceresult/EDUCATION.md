# Once With Result

## Intuition

`sync.Once` guarantees the function runs once, not that it succeeds. Anything the function can do — return an error, panic — must be turned into cached state, or later callers get behaviour that depends on who arrived first.

## Approach

1. Wrap the body in `o.once.Do`.
2. Increment `runs` first so it reflects real executions.
3. Defer a `recover` that converts a panic into `o.err`.
4. Assign both results inside the once function; return the cached fields afterwards.

## Solution

```go
func (o *OnceValue) Get() (string, error) {
	o.once.Do(func() {
		o.runs++
		defer func() {
			if r := recover(); r != nil {
				o.value = ""
				o.err = fmt.Errorf("init panicked: %v", r)
			}
		}()
		o.value, o.err = o.init.Init()
	})
	return o.value, o.err
}
```

## Walkthrough

Without the `recover`, the panic unwinds through `Do`. `Once` still marks itself done, so every later `Get` returns a zero value with a nil error — a silent, permanent corruption instead of a visible failure.

## Pitfalls

- Letting the panic escape: `Once` records completion regardless, so the failure disappears after the first call.
- Retrying on error by resetting the `Once` — a `sync.Once` cannot be reset, and copying it is a race.
- Returning the fields from outside `Do` without any synchronisation — the `Do` call is what publishes them.
