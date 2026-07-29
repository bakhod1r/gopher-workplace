# Map values are not addressable

## Intuition

A map returns a **copy** of its value, and map elements aren't addressable — so
`m[key].Hits++` won't even compile. The pattern is read-modify-write:

```go
s := m[key]
s.Hits++
m[key] = s
```

## Approach

1. Bug: s := m[key] copies the struct value; s.Hits++ mutates the copy and _ = s throws it away, so the map is unchanged. 2. Map values are not addressable, so you must write the modified struct back. 3. Fix: m[key] = s.

## Solution

```go
type Stat struct {
	Hits int
}

func Record(m map[string]Stat, key string) {
	s := m[key]
	s.Hits++
	m[key] = s
}
```

## Walkthrough

s=m[key] copies {Hits:0}; s.Hits++ -> {Hits:1} in the copy; without write-back m[key] stays {Hits:0}. m[key]=s stores the incremented struct.

## Pitfalls

- `m[k]` yields a copy; mutating it doesn't touch the map.
- `m[k].Field = x` and `m[k].Field++` are compile errors (unaddressable).
- Pointer values (`map[K]*V`) are mutable in place.
