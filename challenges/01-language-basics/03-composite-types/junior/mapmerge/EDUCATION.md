# Copying and merging maps

## Intuition

Maps are reference types, so returning a fresh, independent map means allocating
one and copying entries in override order:

```go
out := make(map[string]int)
for k, v := range a { out[k] = v }
for k, v := range b { out[k] = v } // b wins on collisions
```

## Approach

1. make a fresh result map.
2. Copy all of a into it.
3. Copy all of b, overwriting on collisions so b wins.
4. Return result; a and b are never modified.

## Solution

```go
func Merge(a, b map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range a {
		result[k] = v
	}
	for k, v := range b {
		result[k] = v
	}
	return result
}
```

## Walkthrough

Merge({"x":1,"y":2},{"y":20,"z":3}): copy a -> {x:1,y:2}; copy b -> y overwritten to 20, z added -> {x:1,y:20,z:3}.

## Pitfalls

- Assigning one map to another copies the **reference**, not the data.
- Iteration order over a map is randomized — don't rely on it.
- Writing to a nil map panics; `make` first.
