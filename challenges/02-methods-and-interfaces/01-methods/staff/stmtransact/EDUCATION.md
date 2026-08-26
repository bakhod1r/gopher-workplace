# STM

## Solution

```go
func Tx(tv *TVar, fn func(int) int) {
	tv.val = fn(tv.val)
}
```
