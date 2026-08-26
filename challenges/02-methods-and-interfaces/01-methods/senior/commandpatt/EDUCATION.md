# Command Pattern (Functional)

## Solution

```go
func (inv *Invoker) ExecuteAll() {
	for _, c := range inv.commands {
		c()
	}
	inv.commands = nil
}
```
