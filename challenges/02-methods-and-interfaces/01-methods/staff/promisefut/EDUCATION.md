# Futures

## Solution

```go
func (f *Future) Complete(val int) {
	f.ch <- val
	close(f.ch)
}

func (f *Future) Get() int {
	return <-f.ch
}
```
