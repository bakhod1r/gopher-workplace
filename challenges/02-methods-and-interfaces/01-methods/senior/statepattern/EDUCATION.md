# State Transitions

## Solution

```go
func (d *Document) Publish() {
	switch d.State {
	case Draft:
		d.State = Moderation
	case Moderation:
		d.State = Published
	}
}
```
