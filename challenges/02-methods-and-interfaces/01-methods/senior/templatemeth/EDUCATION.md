# Template Method

## Intuition

Classic template method uses inheritance: a base class calls abstract methods
that subclasses override. Go has no inheritance, so the same idea arrives as
composition — the "base" holds an interface and calls through it. The result is
better, because the varying part is an explicit, swappable field.

## Approach

1. Keep the algorithm — order, separators, error handling — inside `Run`.
2. Call each varying step through the interface.

## Solution

```go
func (t *Template) Run() string {
	return t.impl.DoStep1() + "-" + t.impl.DoStep2()
}
```

## Walkthrough

`&Template{impl: MyTask{}}` stores a `MyTask` value in an interface field. Both
`DoStep1` and `DoStep2` have value receivers, so the value form satisfies
`Step`. `Run` calls them in source order and joins the results: `"a" + "-" +
"b"`.

## Pitfalls

- **A nil `impl`.** Calling a method on a nil interface panics; a constructor
  should require the implementation.
- **Letting the implementation control ordering.** Then it is strategy, not
  template method — the whole point is that `Run` is the authority.
- **Pointer receivers on the steps.** `MyTask{}` would stop satisfying `Step`
  and the composite literal would fail to compile.
