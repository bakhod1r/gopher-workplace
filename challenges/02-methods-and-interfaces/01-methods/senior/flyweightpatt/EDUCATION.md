# Flyweight Pattern

## Solution

```go
func (f *FlyweightFactory) Get(name string) *FontData {
	if font, ok := f.fonts[name]; ok {
		return font
	}
	font := &FontData{data: name}
	f.fonts[name] = font
	return font
}
```
