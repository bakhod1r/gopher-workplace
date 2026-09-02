# Select Semantics

## Intuition

`select` chooses among *ready* cases. A closed channel is permanently ready, so leaving it in the select spins; setting the variable to nil makes that case block forever, which removes it from consideration.

## Approach

1. `TryRecv` is a `select` with a `default` branch.
2. `Drain` loops while either channel is non-nil, nils out a channel on `!ok`, and appends otherwise.
3. Sort at the end so the result does not depend on scheduling.
4. `FirstReady` uses the same nil-out pattern but returns on the first real value.

## Solution

```go
func TryRecv(ch <-chan int) (int, bool, bool) {
	select {
	case v, ok := <-ch:
		return v, ok, true
	default:
		return 0, false, false
	}
}

func Drain(a, b <-chan int) []int {
	var out []int
	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil
				continue
			}
			out = append(out, v)
		case v, ok := <-b:
			if !ok {
				b = nil
				continue
			}
			out = append(out, v)
		}
	}
	sort.Ints(out)
	return out
}
```

## Walkthrough

Without the nil-out, a closed `a` would win the select almost every time — an infinite spin delivering zero values. Nilling it leaves only `b`'s case live, and the loop exits when both are nil.

## Pitfalls

- Adding a `default` to `Drain`'s select, which turns a blocking multiplexer into a busy loop.
- Breaking out of the loop when one channel closes, which loses the other's remaining values.
- Assuming `select` is fair or ordered: it picks uniformly at random among ready cases, which is why the result is sorted.
