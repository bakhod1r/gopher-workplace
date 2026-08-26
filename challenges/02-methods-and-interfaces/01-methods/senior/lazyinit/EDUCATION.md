# Lazy Initialization

## Solution

```go
func (l *LazyString) String() string {
	if l.val == nil {
		v := l.init()
		l.val = &v
	}
	return *l.val
}
```
