# Event Bus

## Solution

```go
func (b *Bus) On(eventType string, listener func(data string)) {
	b.listeners[eventType] = append(b.listeners[eventType], listener)
}

func (b *Bus) Emit(eventType string, data string) {
	for _, l := range b.listeners[eventType] {
		l(data)
	}
}
```
