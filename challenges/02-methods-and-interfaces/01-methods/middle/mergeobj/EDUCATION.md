# Merging with Zero-Value Guards

## Intuition

Go's zero values (`""`, `0`, `false`) serve as "not set" markers. A merge
function skips zero-value fields, applying only explicitly set overrides.

## Approach

1. Check each field: if non-zero, overwrite.

## Solution

```go
func (c *Config) Merge(other Config) {
	if other.Host != "" {
		c.Host = other.Host
	}
	if other.Port != 0 {
		c.Port = other.Port
	}
	if other.Debug {
		c.Debug = other.Debug
	}
}
```

## Walkthrough

`Config{"localhost", 8080, false}.Merge(Config{Port: 9090})`:
- `other.Host` = "" → skip.
- `other.Port` = 9090 ≠ 0 → `c.Port = 9090`.
- `other.Debug` = false → skip.

## Pitfalls

- `false` is the zero value for bool — you can't distinguish "set to false" from
  "not set". This is a known limitation of zero-value merging.
- Using `reflect` would be more general but is overkill here.
