# VM Opcode

## Solution

```go
func (v *VM) Next() int {
	v.IP++
	return v.IP
}
```
