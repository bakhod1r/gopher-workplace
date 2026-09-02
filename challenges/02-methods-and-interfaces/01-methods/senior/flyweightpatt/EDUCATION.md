# Flyweight Pattern

## Intuition

A flyweight is caching with intent: the objects are immutable and
interchangeable, so there is no reason for two of them to exist. The factory
owns the map and is the only way to get an instance, which is what makes the
"one per key" guarantee hold.

## Approach

1. Check the map first.
2. On a miss, build the heavy object once.
3. Store it before returning, so the next call hits.

## Solution

```go
func (f *FlyweightFactory) Get(name string) *FontData {
	if fd, ok := f.fonts[name]; ok {
		return fd
	}
	fd := &FontData{data: name}
	f.fonts[name] = fd
	return fd
}
```

## Walkthrough

`f1 := f.Get("Arial")` misses, allocates, stores, returns the pointer.
`f2 := f.Get("Arial")` hits and returns the *same* pointer, so `f1 == f2` is
true — pointer comparison compares addresses, not contents.
`f.Get("Times")` misses again and allocates a second object, so `f1 != f3`.

## Pitfalls

- **Building before checking.** Allocating first and only then consulting the map
  wastes the work the pattern exists to avoid.
- **Forgetting to store.** Every call then returns a fresh pointer and the
  identity test fails.
- **A nil map.** `f.fonts` must be initialized (the test does it in the literal);
  writing to a nil map panics with `assignment to entry in nil map`.
- **Comparing structs instead of pointers.** `*f1 == *f3` compares contents and
  would hide the bug.
