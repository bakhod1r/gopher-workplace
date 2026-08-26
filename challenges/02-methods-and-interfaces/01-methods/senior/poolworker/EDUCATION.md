# Worker Pool Methods

## Intuition

A worker pool struct encapsulates the channels and wait groups needed to manage
the workers. The `Start()` method is a common pattern to initialize the background
routines.

## Solution

```go
func (p *Pool) Start() {
	for i := 0; i < p.Count; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for task := range p.Tasks {
				task()
			}
		}()
	}
}
```
