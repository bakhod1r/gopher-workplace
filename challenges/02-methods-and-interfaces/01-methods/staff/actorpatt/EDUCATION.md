# Actor Pattern

## Solution

```go
func (a *CounterActor) Add(n int) {
	a.msgs <- func(c *int) {
		*c += n
	}
}
```
