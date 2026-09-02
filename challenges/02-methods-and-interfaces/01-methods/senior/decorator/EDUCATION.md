# Decorator Pattern

## Intuition

Wrapping is a sandwich: do something before, call the inner thing, do something
after. Here "before" and "after" are the two brackets. The wrapped `Component`
never learns it was decorated, which is why decorators stack.

## Approach

1. Ask the inner component for its result.
2. Add the outer behaviour to that result.
3. Return the combined value.

## Solution

```go
func (d *Decorator) DoWork() string {
	return "[" + d.Comp.DoWork() + "]"
}
```

## Walkthrough

`d.Comp` is a `*Component`. `d.Comp.DoWork()` returns `"work"`. Concatenation
produces `"[work]"`. Because the decorator's method has the same *shape* as the
component's — no arguments, one `string` back — a decorator can itself be
wrapped by another decorator, giving `"[[work]]"`.

## Pitfalls

- **Returning `"[work]"` literally.** The test passes; the pattern is gone, and
  it breaks the moment the component's output changes.
- **Nil `Comp`.** A decorator constructed without a component panics on the
  delegating call — real code guards or requires it in a constructor.
- **Embedding instead of a field.** Embedding `*Component` would promote
  `DoWork` and make the outer definition ambiguous to reason about; an explicit
  field keeps the delegation visible.
