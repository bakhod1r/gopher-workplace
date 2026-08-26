# GC Mark

## Solution

```go
func (o *Object) Mark() {
	if o.Marked { return }
	o.Marked = true
	for _, ref := range o.Refs {
		ref.Mark()
	}
}
```
