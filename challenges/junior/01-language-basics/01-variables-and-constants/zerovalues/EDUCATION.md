# Zero values

## The idea

Go has no uninitialised memory. Every declaration without an explicit value gets
its type's **zero value**, and that value is always usable:

| Type | Zero value |
|------|-----------|
| `int`, `float64`, all numerics | `0` |
| `string` | `""` |
| `bool` | `false` |
| pointer, slice, map, channel, func, interface | `nil` |
| struct | every field set to *its* zero value |
| array | every element set to its zero value |

```go
var n int        // 0
var s string     // ""
var c Config     // Config{Host: "", Port: 0, Tags: nil}
```

This is why Go has no "undefined" and no constructor requirement.

## Useful zero values are a design goal

The standard library leans on this. A `sync.Mutex` is ready to lock at zero; a
`bytes.Buffer` is ready to write at zero. When you design a struct, ask what its
zero value should mean, and pick field types so that "empty" is correct rather
than broken.

```go
type Config struct {
	Host  string
	Port  int
	Tags  []string
}

var c Config      // usable: empty host, port 0, no tags
c.Tags = append(c.Tags, "beta")   // append to a nil slice works
```

## `nil` slice vs empty slice

Both have length 0, and both work with `len`, `range` and `append`. They are not
identical:

```go
var a []string        // nil    — len 0, cap 0, a == nil
b := []string{}       // empty  — len 0, but b != nil
```

The difference is visible where identity matters: `a == nil` is true and `b ==
nil` is false, and `encoding/json` marshals the first as `null` and the second
as `[]`. Prefer the nil form unless you specifically need the other — it is the
zero value, and it costs no allocation.

## Declaration forms

```go
var c Config              // all-zero struct
c := Config{}             // same thing, written as a literal
c := Config{Port: 8080}   // named field set; Host and Tags stay zero
```

A composite literal fills the fields you name and zeroes the rest, so partially
configured structs need no boilerplate.

## Watch out

- A `nil` map is readable but **not writable**: `m["k"]` returns the zero value,
  `m["k"] = 1` panics. Maps need `make` before writing; slices do not need it
  before `append`.
- A `nil` pointer dereference panics. A zero *struct* is safe; a zero *pointer to
  struct* is not.
- `var s []string` and `s := []string(nil)` are the same thing written twice.
- Zero values are set for you — writing `var n int = 0` adds nothing.

## Try it yourself

```go
var (
	n int
	s string
	b bool
	sl []int
	m map[string]int
)
fmt.Println(n, len(s), b, sl == nil, m == nil)   // 0 0 false true true
fmt.Println(len(sl), len(m), m["missing"])       // 0 0 0
sl = append(sl, 1)                               // fine on a nil slice
// m["k"] = 1                                    // panic: assignment to entry in nil map
```
