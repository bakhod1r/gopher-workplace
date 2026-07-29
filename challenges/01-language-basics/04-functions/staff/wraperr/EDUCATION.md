# The defer-wrap-error idiom

## Intuition

The idiomatic error decorator captures the NAMED return in the closure body; passing it as a deferred argument freezes the pre-assignment (nil) value.

## Approach

1. A deferred error-wrapper must read the **named return** at return time.
2. The bug snapshots `err` as a deferred argument (nil when defer runs).
3. Close over `err` with no parameter.

## Solution

```go
import (
	"errors"
	"fmt"
)

func Do(flag bool) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("do: %w", err)
		}
	}()
	if flag {
		err = errors.New("boom")
	}
	return
}
```

## Walkthrough

`defer func(e error){...}(err)` captures `err` while it is still nil, so nothing wraps. Reading the named `err` inside a parameterless defer sees the final error and wraps it.

## Pitfalls

- `defer func(e error){...}(err)` snapshots err (nil) at defer-time.
- Read the named `err` in the body so you see the final value.
