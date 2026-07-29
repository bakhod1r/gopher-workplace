# Reverse lookups

## Intuition

Inverting a map swaps roles: each `key→value` becomes `value→key`. The result
map's types are the originals reversed:

```go
out := make(map[int]string)
for k, v := range m { out[v] = k }
```

## Approach

1. make an empty map[int]string.
2. Range the input as k, v.
3. Store result[v] = k (values become keys).
4. Return the inverted map.

## Solution

```go
func Invert(m map[string]int) map[int]string {
	result := make(map[int]string)
	for k, v := range m {
		result[v] = k
	}
	return result
}
```

## Walkthrough

Invert({"one":1,"two":2}): store 1->"one", 2->"two". Result {1:"one",2:"two"}.

## Pitfalls

- If values are **not** unique, later pairs overwrite earlier ones — inversion
  loses data.
- The value type must be a valid map key (comparable).
- Iteration order is random; the resulting map is unordered.
