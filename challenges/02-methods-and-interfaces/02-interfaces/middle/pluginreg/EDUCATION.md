# Plugin Registry

## Intuition

A small required interface plus optional ones keeps the contract honest: plugins implement only what they need, and the host discovers extras with an assertion.

## Approach

1. `Register` checks the name set, records it, and appends to the ordered slice.
2. `RunAll` collects `Run()` output in slice order.
3. `CloseAll` asserts each plugin to `Closer` and closes the ones that match.

## Solution

```go
func (r *Registry) Register(p Plugin) error {
	if r.names[p.Name()] {
		return ErrDuplicate
	}
	r.names[p.Name()] = true
	r.plugins = append(r.plugins, p)
	return nil
}

func (r *Registry) RunAll() []string {
	out := make([]string, 0, len(r.plugins))
	for _, p := range r.plugins {
		out = append(out, p.Run())
	}
	return out
}

func (r *Registry) CloseAll() int {
	n := 0
	for _, p := range r.plugins {
		if c, ok := p.(Closer); ok {
			c.Close()
			n++
		}
	}
	return n
}
```

## Walkthrough

`Simple` has no `Close`, so `p.(Closer)` fails and it is skipped; `Closeable` matches, is closed, and counted — hence 1 out of 2.

## Pitfalls

- Storing plugins only in a map, which loses registration order.
- Appending before the duplicate check, so a rejected plugin still runs.
- Requiring `Close` on the main `Plugin` interface, which forces every plugin to write an empty method.
