# Pub-Sub with RWMutex

## Solution

```go
func (ps *PubSub) Publish(topic, msg string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	for _, ch := range ps.subs[topic] {
		ch <- msg
	}
}
```
