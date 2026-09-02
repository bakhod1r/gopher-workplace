# Interface Call Cost

## Intuition

An interface call loads a function pointer from the itab and jumps through it. The call itself is a few instructions, but the indirection blocks inlining — so the real cost is the optimisations that do not happen inside the loop.

## Approach

1. Both `Apply` methods are one-liners over value receivers.
2. `RunIface` and `RunConcrete` are the same fold; only the parameter type differs.
3. Neither allocates: an `AddOp` stored in an interface at the call site is boxed once, outside the measured loop.
4. Compare the two benchmarks to see the dispatch cost.

## Solution

```go
func (o AddOp) Apply(acc, v int) int { return acc + v + o.N }

func (MulOp) Apply(acc, v int) int { return acc * v }

func RunIface(op Op, start int, vs []int) int {
	acc := start
	for _, v := range vs {
		acc = op.Apply(acc, v)
	}
	return acc
}

func RunConcrete(op AddOp, start int, vs []int) int {
	acc := start
	for _, v := range vs {
		acc = op.Apply(acc, v)
	}
	return acc
}
```

## Walkthrough

Run `go build -gcflags='-m -m'` on this package: the concrete `Apply` is reported as inlined into `RunConcrete`, while the interface call in `RunIface` is not, unless the compiler can prove the dynamic type.

## Pitfalls

- Concluding "interfaces are slow" — the dispatch is cheap; the lost inlining and boxing are what cost.
- Micro-benchmarking with a loop the compiler can delete: always consume the result.
- Refactoring to concrete types before profiling; this cost matters only in genuinely hot loops.
