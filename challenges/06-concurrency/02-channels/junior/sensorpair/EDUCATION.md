# select With Two Ready Channels

## Intuition

`select` is a switch over channel operations. It blocks until at least one
case can proceed; if several can, it picks *uniformly at random*. That
randomness prevents starvation, but it means you must never rely on case
order. Here each sensor sends exactly one reading, so two iterations drain
both regardless of the order chosen.

## Approach

1. Start `total` at 0.
2. Loop exactly twice.
3. In each iteration, `select` on `<-temp` and `<-humidity`, adding whichever fires.
4. Return `total`.

## Solution

```go
func CombinedReading(temp, humidity <-chan int) int {
	total := 0
	for i := 0; i < 2; i++ {
		select {
		case v := <-temp:
			total += v
		case v := <-humidity:
			total += v
		}
	}
	return total
}
```

## Walkthrough

With `temp` holding `21` and `humidity` holding `40`: iteration one might
take either; say it takes `40` (total 40). Iteration two finds only `temp`
ready and takes `21` (total 61). The reverse order gives the same `61`.

## Pitfalls

- Reading `temp` twice would happen if it were buffered with two readings — the loop counts *receives*, not channels.
- A `select` with no ready case and no `default` blocks; a silent sensor deadlocks the dashboard.
- Case order in the source has no effect on which case is chosen.
