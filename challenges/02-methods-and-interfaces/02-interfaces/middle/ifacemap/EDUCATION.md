# Interface Map

## Intuition

A `map[string]Command` turns dispatch into data. New commands are registered at runtime, and the dispatcher never changes.

## Approach

1. `Register` writes into the map, so re-registering replaces.
2. `Run` uses the comma-ok lookup and returns `ErrUnknown` when absent.
3. `Names` collects keys into a preallocated slice and sorts them.

## Solution

```go
func (r *Registry) Register(name string, c Command) { r.cmds[name] = c }

func (r *Registry) Run(name, arg string) (string, error) {
	c, ok := r.cmds[name]
	if !ok {
		return "", ErrUnknown
	}
	return c.Exec(arg), nil
}

func (r *Registry) Names() []string {
	out := make([]string, 0, len(r.cmds))
	for k := range r.cmds {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
```

## Walkthrough

Registering `"x"` twice leaves one map entry, and the second command wins — so `Run("x", "AB")` lowercases and `Names` has length 1.

## Pitfalls

- Reading `r.cmds[name]` without `ok`: a missing key yields a nil `Command`, and calling `Exec` on it panics.
- Returning `Names` unsorted, which flakes because map order is randomised.
- Building `Registry{}` directly, leaving `cmds` nil so `Register` panics.
