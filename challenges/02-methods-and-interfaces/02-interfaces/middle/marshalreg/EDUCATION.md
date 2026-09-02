# Codec Registry

## Intuition

Resolution order is the design: exact match wins, then a configured default, and only then an error. Storing the default as an interface field makes "unset" representable as nil.

## Approach

1. Both codecs delegate to `strings.Join` with their separator.
2. `Register` and `SetDefault` are plain assignments.
3. `Encode` looks the type up with comma-ok, then checks `r.fallbck == nil`, then encodes with the fallback.

## Solution

```go
func (CSVCodec) Encode(fields []string) string { return strings.Join(fields, ",") }

func (PipeCodec) Encode(fields []string) string { return strings.Join(fields, "|") }

func (r *Registry) Register(contentType string, c Codec) { r.codecs[contentType] = c }

func (r *Registry) SetDefault(c Codec) { r.fallbck = c }

func (r *Registry) Encode(contentType string, fields []string) (string, error) {
	if c, ok := r.codecs[contentType]; ok {
		return c.Encode(fields), nil
	}
	if r.fallbck == nil {
		return "", ErrNoCodec
	}
	return r.fallbck.Encode(fields), nil
}
```

## Walkthrough

With no `SetDefault`, `r.fallbck` is the nil interface, so `Encode("unknown", ...)` returns `ErrNoCodec` instead of panicking on a nil method call.

## Pitfalls

- Calling `r.fallbck.Encode(...)` without the nil check — a nil interface call panics.
- Returning `ErrNoCodec` before consulting the fallback.
- Using the zero `Registry`, whose `codecs` map is nil, so `Register` panics.
