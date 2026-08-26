# Template Method

## Solution

```go
func (t *Template) Run() string {
	return t.impl.DoStep1() + "-" + t.impl.DoStep2()
}
```
