# Custom Formatter

## Solution

```go
func (n Name) Format() string { return n.Last + ", " + n.First }
```
