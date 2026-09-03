# The Frames That Are Not Really There

## Intuition

The profiler samples a program counter. That address belongs to exactly one real function; every other name on the stack came out of the compiler's inline tree.

## Approach

1. `Physical` filters on `!Inlined`.
2. `Attribute` walks from the leaf backwards to the first non-inlined frame.

## Solution

```go
func Physical(stack []Frame) []string {
	out := make([]string, 0, len(stack))
	for _, f := range stack {
		if !f.Inlined {
			out = append(out, f.Func)
		}
	}
	return out
}

func Attribute(stack []Frame) (string, bool) {
	for i := len(stack) - 1; i >= 0; i-- {
		if !stack[i].Inlined {
			return stack[i].Func, true
		}
	}
	return "", false
}
```

## Walkthrough

For `[{a false} {b true} {c false}]` the leaf `c` is real, so it keeps the attribution; had `c` been inlined, the cost would belong to `a`, the nearest function with actual machine code.

## Pitfalls

- Scanning forward and returning the *outermost* physical frame, which blames `main` for everything.
- Assuming an inlined frame is free — its work is real, it just lives in the caller's body.
- Comparing profiles across builds without checking inlining decisions changed underneath you.
