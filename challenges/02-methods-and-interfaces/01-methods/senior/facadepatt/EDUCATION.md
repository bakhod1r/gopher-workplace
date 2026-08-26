# Facade Pattern

## Solution

```go
func (f *Facade) Operation() string {
	return f.s1.Op1() + "+" + f.s2.Op2()
}
```
