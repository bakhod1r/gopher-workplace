# iota with shifts: binary units

## Intuition

Storage sizes step by factors of 1024 (2^10). One `iota` expression generates
the whole ladder:

```go
type ByteSize uint64
const (
	_  ByteSize = iota      // skip iota==0 with the blank identifier
	KB ByteSize = 1 << (10 * iota) // iota==1 -> 1<<10 == 1024
	MB                              // iota==2 -> 1<<20
	GB                              // iota==3 -> 1<<30
	TB                              // iota==4 -> 1<<40
)
```

The blank identifier `_` consumes the `iota==0` slot (where `1<<0` would be 1),
so `KB` starts at the first useful power.

## Approach

1. Define KB..TB with `1 << (10 * iota)`, skipping slot 0.
2. In `Humanize`, try units largest-first.
3. Return the first that divides `n` evenly; else bytes.

## Solution

```go
type ByteSize uint64

const (
	_  ByteSize = iota
	KB          = 1 << (10 * iota)
	MB
	GB
	TB
)

func Humanize(n ByteSize) (ByteSize, string) {
	units := []struct {
		v ByteSize
		s string
	}{{TB, "TB"}, {GB, "GB"}, {MB, "MB"}, {KB, "KB"}}
	for _, u := range units {
		if n >= u.v && n%u.v == 0 {
			return n / u.v, u.s
		}
	}
	return n, "B"
}
```

## Walkthrough

`Humanize(1048576)` checks TB, GB (no), then MB — `1048576 % MB == 0` → `1, "MB"`.

## Pitfalls

- `uint64` (not `int`) because `1 << 40` overflows a 32-bit int.
- The `_` line is what shifts KB onto `iota==1`; without it KB would be 1.
- Binary KB = 1024. If you need decimal (1000), that is `kB`, a different unit.
