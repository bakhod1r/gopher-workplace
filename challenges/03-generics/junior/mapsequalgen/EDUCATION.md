# Compare Maps

## Intuition

The constraint difference between `Clone` and `Equal` is not arbitrary — copying a value needs nothing, comparing one needs `==`.

## Approach

1. Return `maps.Equal(a, b)`.

## Solution

```go
func SameConfig[K, V comparable](a, b map[K]V) bool {
	return maps.Equal(a, b)
}
```

## Walkthrough

`SameConfig(nil, {})` sees two maps of size 0 with no differing keys, so it returns `true`.

## Pitfalls

- Declaring `V any`, which cannot be compared.
- Comparing only lengths.
- Using `reflect.DeepEqual` where the typed helper is available.
