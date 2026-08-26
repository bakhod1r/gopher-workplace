# sync.Once Initialization

## Solution

```go
func (l *LazyData) Get() string {
	l.once.Do(func() {
		l.data = l.init()
	})
	return l.data
}
```
