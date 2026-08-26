# Memento Pattern

## Solution

```go
func (e *Editor) Save() Memento {
	return Memento{state: e.Text}
}

func (e *Editor) Restore(m Memento) {
	e.Text = m.state
}
```
