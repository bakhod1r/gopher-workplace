# Sync Pool

## Solution

```go
func (p *BufferPool) Get() *Buffer {
	return p.pool.Get().(*Buffer)
}

func (p *BufferPool) Put(b *Buffer) {
	p.pool.Put(b)
}
```
