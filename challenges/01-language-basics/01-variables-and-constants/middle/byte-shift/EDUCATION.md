# iota with shifts: binary units

## The idea

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

## Why it matters

These constants are folded at compile time — `2 * KB` costs nothing at runtime.
They also document intent far better than `2097152`.

## Watch out

- `uint64` (not `int`) because `1 << 40` overflows a 32-bit int.
- The `_` line is what shifts KB onto `iota==1`; without it KB would be 1.
- Binary KB = 1024. If you need decimal (1000), that is `kB`, a different unit.

## Try it yourself

```go
const (
	_   = iota
	KHz = 1 << (10 * iota) // not physical, just to see the pattern
	MHz
)
```
