# Builder Pattern

## Solution

```go
func (b *RequestBuilder) URL(u string) *RequestBuilder {
	b.req.URL = u
	return b
}

func (b *RequestBuilder) Auth(t string) *RequestBuilder {
	b.req.Auth = t
	return b
}
```
